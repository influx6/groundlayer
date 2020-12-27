package peji

import (
	"bytes"
	"context"
	"strings"
	"sync"
	"time"

	"github.com/influx6/npkg/njson"
	"github.com/influx6/sabuhp"
	"github.com/influx6/sabuhp/mixer"

	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nxid"

	"github.com/influx6/groundlayer/pkg/domu"
)

const (
	DefaultMaxPageIdleness            = 5 * time.Minute
	DefaultPageIdlenessChecksInterval = 2 * time.Minute

	HeaderSessionIdName         = "X-Void-Id"
	QueryAndCookieSessionIdName = "_groundlayer_id"
)

type PageCreator func(name string, pubsub sabuhp.Transport) *Page

type PageSession struct {
	lastUsed time.Time
	Id       nxid.ID
	Page     *Page
}

func (ps *PageSession) Close() {
	ps.lastUsed = time.Time{}
	ps.Page.Close()
}

type Logger interface {
	Log(json *njson.JSON)
}
type PageSessionManager struct {
	routePath         string
	maxIdle           time.Duration
	idleCheckInterval time.Duration
	onAddRoute        OnPage
	Creator           PageCreator
	sessions          map[string]PageSession
	doFunc            chan func()
	canceler          context.CancelFunc
	ctx               context.Context
	waiter            sync.WaitGroup
	starter           sync.Once
	ender             sync.Once
	rl                sync.Mutex
	routes            map[string]bool
}

func NewPageSessionManager(
	ctx context.Context,
	routePath string,
	maxIdle time.Duration,
	idleCheckInterval time.Duration,
	creator PageCreator,
	onAddRoute OnPage,
) *PageSessionManager {
	var newCtx, canceler = context.WithCancel(ctx)
	return &PageSessionManager{
		routePath:         routePath,
		ctx:               newCtx,
		canceler:          canceler,
		maxIdle:           maxIdle,
		Creator:           creator,
		onAddRoute:        onAddRoute,
		idleCheckInterval: idleCheckInterval,
		doFunc:            make(chan func(), 0),
		routes:            map[string]bool{},
		sessions:          map[string]PageSession{},
	}
}

func (psm *PageSessionManager) GetName() string {
	return psm.routePath
}

func (psm *PageSessionManager) Wait() {
	psm.waiter.Wait()
}

func (psm *PageSessionManager) Stop() {
	psm.starter.Do(func() {
		psm.canceler()
		psm.waiter.Wait()
	})
}

func (psm *PageSessionManager) Start() {
	psm.starter.Do(func() {
		psm.waiter.Add(1)
		go psm.manage()
	})
}

type SessionStat struct {
	PageName      string
	TotalSessions int
	Sessions      map[string]time.Time
}

func (psm *PageSessionManager) Stat() (SessionStat, error) {
	var session = make(chan SessionStat, 1)

	psm.doFunc <- func() {
		var stat SessionStat
		stat.PageName = psm.routePath
		stat.Sessions = map[string]time.Time{}
		stat.TotalSessions = len(psm.sessions)

		for _, ss := range psm.sessions {
			stat.Sessions[ss.Id.String()] = ss.lastUsed
		}

		session <- stat
	}

	select {
	case <-psm.ctx.Done():
		return SessionStat{}, nerror.WrapOnly(psm.ctx.Err())
	case ss := <-session:
		return ss, nil
	}
}

// Retire returns closes a specific session using giving id.
func (psm *PageSessionManager) Retire(sessionId string) error {
	var session = make(chan error, 1)

	psm.doFunc <- func() {
		if ss, hasSession := psm.sessions[sessionId]; hasSession {
			delete(psm.sessions, sessionId)
			ss.Page.Close()
		}
		session <- nil
	}

	select {
	case <-psm.ctx.Done():
		return nerror.WrapOnly(psm.ctx.Err())
	case <-session:
		return nil
	}
}

// Session returns a giving page for a giving sessionId.
func (psm *PageSessionManager) Session(sessionId string) (*Page, error) {
	var session = make(chan *Page, 1)
	var errs = make(chan error, 1)

	psm.doFunc <- func() {
		var ss, hasSession = psm.sessions[sessionId]
		if !hasSession {
			errs <- nil
			return
		}
		session <- ss.Page
	}

	select {
	case <-psm.ctx.Done():
		return nil, nerror.WrapOnly(psm.ctx.Err())
	case err := <-errs:
		return nil, nerror.WrapOnly(err)
	case page := <-session:
		return page, nil
	}
}

// NewSession returns a new session page and session id.
func (psm *PageSessionManager) NewSession(t sabuhp.Transport) (*Page, string, error) {
	var session = make(chan PageSession, 1)

	psm.doFunc <- func() {
		var ps PageSession
		ps.Id = nxid.New()
		ps.lastUsed = time.Now()
		ps.Page = psm.Creator(psm.routePath, t)
		psm.sessions[ps.Id.String()] = ps

		ps.Page.OnPageAdd(psm.manageAddPageRoute)
		session <- ps
	}

	select {
	case <-psm.ctx.Done():
		return nil, "", nerror.WrapOnly(psm.ctx.Err())
	case ss := <-session:
		return ss.Page, ss.Id.String(), nil
	}
}

func (psm *PageSessionManager) manageAddPageRoute(pageRoute string, p *Page) {
	psm.rl.Lock()
	if _, hasRoute := psm.routes[pageRoute]; hasRoute {
		psm.rl.Unlock()
		return
	}
	psm.routes[pageRoute] = true
	psm.rl.Unlock()

	psm.onAddRoute(pageRoute, p)
}

func (psm *PageSessionManager) manage() {
	defer psm.waiter.Done()

	var ticker = time.NewTicker(psm.idleCheckInterval)
	defer ticker.Stop()

doLoop:
	for {
		select {
		case <-psm.ctx.Done():
			return
		case doFn := <-psm.doFunc:
			doFn()
		case <-ticker.C:
			// clean house
			var nowTime = time.Now()
			for key, session := range psm.sessions {
				if nowTime.Sub(session.lastUsed) < psm.maxIdle {
					continue doLoop
				}
				delete(psm.sessions, key)
				session.Close()
			}
		}
	}
}

var _ sabuhp.TransportResponse = (*Pages)(nil)

type OnPages func(route string, p *Pages)

// Pages exists to provider an organization
// around sessions and pages.
//
// It implements the http.Handler interface.
type Pages struct {
	logger    Logger
	prefix    string
	router    *mixer.Mux
	maxIdle   time.Duration
	idleCheck time.Duration
	ctx       context.Context
	tr        sabuhp.Transport
	sl        sync.RWMutex
	waiter    sync.WaitGroup
	managers  map[string]*PageSessionManager
	onNewPage *PageNotification
}

func WithPages(
	ctx context.Context,
	logger Logger,
	prefix string,
	transport sabuhp.Transport,
	notFound Handler,
) *Pages {
	return NewPages(
		ctx,
		logger,
		prefix,
		DefaultMaxPageIdleness,
		DefaultPageIdlenessChecksInterval,
		transport,
		notFound,
	)
}

func NewPages(
	ctx context.Context,
	logger Logger,
	prefix string,
	maxIdle time.Duration,
	idleCheck time.Duration,
	transport sabuhp.Transport,
	notFound Handler,
) *Pages {
	if notFound == nil {
		notFound = DefaultNotFound{}
	}

	return &Pages{
		tr:        transport,
		prefix:    prefix,
		ctx:       ctx,
		logger:    logger,
		maxIdle:   maxIdle,
		idleCheck: idleCheck,
		onNewPage: NewPageNotification(),
		managers:  map[string]*PageSessionManager{},
		router: mixer.NewMux(mixer.MuxConfig{
			RootPath: prefix,
			NotFound: mixer.HandlerFunc(func(message *sabuhp.Message) (*sabuhp.Message, error) {
				var d Data
				d.Message = message
				d.Path = message.Path
				if sessionId, sessionIdErr := getSessionId(message); sessionIdErr == nil {
					d.SessionId = sessionId
				}

				var payload, contentType, err = writeNode(message.ContentType, notFound.Handle(d))
				if err != nil {
					return nil, nerror.WrapOnly(err)
				}

				return &sabuhp.Message{
					MessageMeta: sabuhp.MessageMeta{
						Path:                message.Path,
						ContentType:         contentType,
						SuggestedStatusCode: 200,
						Headers: sabuhp.Header{
							HeaderSessionIdName: []string{d.SessionId},
						},
						Cookies: []sabuhp.Cookie{
							{
								Name:  QueryAndCookieSessionIdName,
								Value: d.SessionId,
							},
						},
					},
					FromAddr: prefix,
					ID:       nxid.ID{},
					Topic:    message.Path,
					Delivery: message.Delivery,
					Payload:  payload,
					Metadata: sabuhp.Params{},
					Params:   sabuhp.Params{},
				}, nil
			}),
		}),
	}
}

func (p *Pages) Wait() {
	p.waiter.Wait()
}

type PagesStat struct {
	TotalPages   int
	PageSessions map[string]SessionStat
}

// Stats returns current states of existing pages, creators.
func (p *Pages) Stats() PagesStat {
	var totalPages int
	var stats = map[string]SessionStat{}

	p.sl.RLock()
	totalPages = len(p.managers)
	for page, manager := range p.managers {
		stat, err := manager.Stat()
		if err != nil {
			var stack = njson.Log(p.logger)
			stack.LError().Error("error", err).
				Message("failed to get stat for page").
				String("page", page).
				End()
			njson.ReleaseLogStack(stack)
			continue
		}
		stats[page] = stat
	}
	p.sl.RUnlock()

	var stat PagesStat
	stat.PageSessions = stats
	stat.TotalPages = totalPages
	return stat
}

// GetManager returns page creator registered for a giving page page
func (p *Pages) Get(pageName string) (*PageSessionManager, error) {
	var pageHandle = cleanAllSlashes(handlePath(pageName))
	var prefixPage = cleanAllSlashes(handlePath(p.prefix, pageName))

	p.sl.RLock()
	var manager, exists = p.managers[prefixPage]
	if !exists {
		manager, exists = p.managers[pageHandle]
	}
	p.sl.RUnlock()

	if !exists {
		return nil, nerror.New("not found")
	}
	return manager, nil
}

// HasPage returns true/false if a giving page exists.
func (p *Pages) Has(pageName string) bool {
	var pageHandle = cleanAllSlashes(handlePath(pageName))
	var prefixPage = cleanAllSlashes(handlePath(p.prefix, pageName))
	p.sl.RLock()
	if _, exists := p.managers[prefixPage]; exists {
		p.sl.RUnlock()
		return true
	}
	if _, exists := p.managers[pageHandle]; exists {
		p.sl.RUnlock()
		return true
	}
	p.sl.RUnlock()
	return false
}

// AddCreator adds a new PageCreator for a giving page routePath.
// It returns true/false based on whether the routePath and creator was registered or if there was routePath conflict.
func (p *Pages) Add(pageName string, creatorFunc PageCreator) error {
	var prefixPage = cleanAllSlashes(handlePath(p.prefix, pageName))

	var routerPath = cleanAllSlashes(handlePath(pageName))
	var routerPathForMore = cleanAllSlashes(handlePath(p.prefix, pageName, "*path"))

	if p.Has(prefixPage) {
		return nerror.New("already exists")
	}

	var manager = NewPageSessionManager(p.ctx, prefixPage, p.maxIdle, p.idleCheck, creatorFunc, p.onNewPage.Emit)
	manager.Start()

	p.waiter.Add(1)
	go func() {
		defer p.waiter.Done()
		manager.Wait()

		p.sl.Lock()
		delete(p.managers, prefixPage)
		p.sl.Unlock()
	}()

	p.sl.Lock()
	p.managers[prefixPage] = manager
	p.sl.Unlock()

	var handler = createHandler(prefixPage, manager, p.tr)
	p.router.Serve(routerPath, handler)
	p.router.Serve(routerPathForMore, handler)
	p.onNewPage.Emit(prefixPage, nil)
	return nil
}

func (p *Pages) AddOnPageRoute(cb OnPages) {
	p.onNewPage.Add(func(route string, _ *Page) {
		cb(route, p)
	})
}

func (p *Pages) Handle(message *sabuhp.Message, tr sabuhp.Transport) sabuhp.MessageErr {
	var reply, err = p.router.ServeRoute(message)
	if err != nil {
		return sabuhp.WrapErr(err, false)
	}

	var sendErr error
	if reply.Delivery == sabuhp.SendToAll {
		sendErr = tr.SendToAll(reply, -1)
	} else {
		sendErr = tr.SendToOne(reply, -1)
	}

	if sendErr != nil {
		return sabuhp.WrapErr(sendErr, false)
	}
	return nil
}

func createHandler(pagePath string, manager *PageSessionManager, tr sabuhp.Transport) mixer.Handler {
	return mixer.HandlerFunc(func(message *sabuhp.Message) (*sabuhp.Message, error) {
		var d Data
		d.Message = message
		d.Path = message.Path
		if sessionId, sessionIdErr := getSessionId(message); sessionIdErr == nil {
			d.SessionId = sessionId
		}

		var page *Page
		var sessionErr error

		page, sessionErr = manager.Session(d.SessionId)
		if sessionErr != nil {
			page, d.SessionId, sessionErr = manager.NewSession(tr)
		}

		if sessionErr != nil {
			return nil, nerror.WrapOnly(sessionErr)
		}

		var renderNode = page.Render(d)
		var payload, contentType, writeErr = writeNode(message.ContentType, renderNode)
		if writeErr != nil {
			return nil, nerror.WrapOnly(writeErr)
		}

		return &sabuhp.Message{
			MessageMeta: sabuhp.MessageMeta{
				Path:                message.Path,
				ContentType:         contentType,
				SuggestedStatusCode: 200,
				Headers: sabuhp.Header{
					HeaderSessionIdName: []string{d.SessionId},
				},
				Cookies: []sabuhp.Cookie{
					{
						Name:  QueryAndCookieSessionIdName,
						Value: d.SessionId,
					},
				},
			},
			ID:       nxid.ID{},
			Topic:    message.Path,
			FromAddr: pagePath,
			Delivery: message.Delivery,
			Payload:  payload,
			Metadata: sabuhp.Params{},
			Params:   sabuhp.Params{},
		}, nil
	})
}

func writeNode(contentType string, renderNode *domu.Node) ([]byte, string, error) {
	var content string
	var renderedOutput = bytes.NewBuffer(make([]byte, 0, 512))
	switch contentType {
	case VoidHTMLDiff:
		content = VoidHTMLDiff
		if renderErr := RenderVoidHTMLDiff(renderNode, renderedOutput); renderErr != nil {
			return nil, "", nerror.WrapOnly(renderErr)
		}
	case VoidJSON:
		content = VoidJSON
		if renderErr := RenderVoidJSON(renderNode, renderedOutput); renderErr != nil {
			return nil, "", nerror.WrapOnly(renderErr)
		}
	case VoidJSONStream:
		content = VoidJSONStream
		if renderErr := RenderVoidJSONStream(renderNode, renderedOutput); renderErr != nil {
			return nil, "", nerror.WrapOnly(renderErr)
		}
	case VoidHTML:
		fallthrough
	default:
		content = PlainHTML
		if renderErr := RenderVoidHTML(renderNode, renderedOutput); renderErr != nil {
			return nil, "", nerror.WrapOnly(renderErr)
		}
	}
	return renderedOutput.Bytes(), content, nil
}

func getSessionId(message *sabuhp.Message) (string, error) {
	var sessionIdFromQuery = strings.TrimSpace(message.Query.Get(QueryAndCookieSessionIdName))
	var sessionIdFromHeader = strings.TrimSpace(message.Headers.Get(HeaderSessionIdName))
	var isASession = len(sessionIdFromHeader) != 0 || len(sessionIdFromQuery) != 0

	if !isASession {
		return "", nerror.New("not a session")
	}
	if len(sessionIdFromQuery) != 0 {
		return sessionIdFromQuery, nil
	}
	return sessionIdFromHeader, nil
}
