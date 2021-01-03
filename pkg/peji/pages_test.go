package peji_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/influx6/sabuhp"
	"github.com/influx6/sabuhp/codecs"
	"github.com/influx6/sabuhp/testingutils"
	"github.com/stretchr/testify/require"

	"github.com/influx6/npkg/njson"

	"github.com/influx6/groundlayer/pkg/peji"
	"github.com/influx6/groundlayer/pkg/styled"
)

type LoggerImpl struct{}

func (l LoggerImpl) Log(json *njson.JSON) {
	log.Println(json.Message())
}

func TestPages(t *testing.T) {
	var newCtx, canceler = context.WithCancel(context.Background())

	var listeners = map[string][]sabuhp.TransportResponse{}
	var transport = &testingutils.TransportImpl{
		ListenFunc: func(topic string, handler sabuhp.TransportResponse) sabuhp.Channel {
			listeners[topic] = append(listeners[topic], handler)
			return &testingutils.NoPubSubChannel{}
		},
		SendToAllFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			return nil
		},
		SendToOneFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			var targetListeners = listeners[data.Topic]
			if len(targetListeners) > 0 {
				_ = targetListeners[0].Handle(data, nil)
			}
			return nil
		},
	}

	var logger = new(LoggerImpl)
	var pages = peji.NewPages(newCtx, logger, "/", 5*time.Second, 2*time.Second, &styled.Theme{}, transport, peji.DefaultNotFound{})

	require.NoError(t, pages.Add("sales", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var sales = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		sales.AddLive("users", &UserComponent{})
		return sales
	}))

	require.NoError(t, pages.Add("about", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var about = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		about.AddLive("users", &UserComponent{})
		return about
	}))

	require.True(t, pages.Has("sales"))
	require.True(t, pages.Has("about"))
	require.False(t, pages.Has("salesforce"))

	var manager, managerErr = pages.Get("sales")
	require.NoError(t, managerErr)

	var page, sessionId, sessionErr = manager.NewSession(nil)
	require.NoError(t, sessionErr)
	require.NotEmpty(t, sessionId)
	require.NotNil(t, page)

	var otherPage, getSessErr = manager.Session(sessionId)
	require.NoError(t, getSessErr)
	require.Equal(t, page, otherPage)

	var pageStat = pages.Stats()
	require.Equal(t, 2, pageStat.TotalPages)

	canceler()
	pages.Wait()
}

func TestPagesWithNotification(t *testing.T) {
	var newCtx, canceler = context.WithCancel(context.Background())

	var listeners = map[string][]sabuhp.TransportResponse{}
	var transport = &testingutils.TransportImpl{
		ListenFunc: func(topic string, handler sabuhp.TransportResponse) sabuhp.Channel {
			listeners[topic] = append(listeners[topic], handler)
			return &testingutils.NoPubSubChannel{}
		},
		SendToAllFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			return nil
		},
		SendToOneFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			var targetListeners = listeners[data.Topic]
			if len(targetListeners) > 0 {
				_ = targetListeners[0].Handle(data, nil)
			}
			return nil
		},
	}

	var logger = new(LoggerImpl)
	var pages = peji.NewPages(newCtx, logger, "/pages/", 5*time.Second, 2*time.Second, &styled.Theme{}, transport, peji.DefaultNotFound{})

	var pageNotification = make(chan string, 2)
	pages.AddOnPageRoute(func(route string, p *peji.Pages) {
		pageNotification <- route
	})

	require.NoError(t, pages.Add("sales", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var sales = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		sales.AddLive("users", &UserComponent{})
		return sales
	}))

	require.NoError(t, pages.Add("about", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var about = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		about.AddLive("users", &UserComponent{})
		return about
	}))

	require.True(t, pages.Has("/pages/sales"))
	require.True(t, pages.Has("/pages/about"))
	require.False(t, pages.Has("/pages/salesforce"))

	var manager, managerErr = pages.Get("/pages/sales")
	require.NoError(t, managerErr)

	var page, sessionId, sessionErr = manager.NewSession(nil)
	require.NoError(t, sessionErr)
	require.NotEmpty(t, sessionId)
	require.NotNil(t, page)

	var otherPage, getSessErr = manager.Session(sessionId)
	require.NoError(t, getSessErr)
	require.Equal(t, page, otherPage)

	var pageStat = pages.Stats()
	require.Equal(t, 2, pageStat.TotalPages)

	var route = <-pageNotification
	require.NotEmpty(t, route)
	require.Equal(t, "/pages/sales", route)

	var route2 = <-pageNotification
	require.NotEmpty(t, route2)
	require.Equal(t, "/pages/about", route2)

	canceler()
	pages.Wait()
}

func TestPagesRouter(t *testing.T) {
	var newCtx, canceler = context.WithCancel(context.Background())

	var logger = new(LoggerImpl)
	var codec = &codecs.JsonCodec{}
	var transposer = sabuhp.NewCodecTransposer(codec, logger, -1)

	var results = make(chan *sabuhp.Message, 1)
	var listeners = map[string][]sabuhp.TransportResponse{}
	var transport = &testingutils.TransportImpl{
		ListenFunc: func(topic string, handler sabuhp.TransportResponse) sabuhp.Channel {
			listeners[topic] = append(listeners[topic], handler)
			return &testingutils.NoPubSubChannel{}
		},
		SendToAllFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			return nil
		},
		SendToOneFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			var targetListeners = listeners[data.Topic]
			if len(targetListeners) > 0 {
				_ = targetListeners[0].Handle(data, nil)
			}
			return nil
		},
	}

	var pages = peji.NewPages(newCtx, logger, "/", 5*time.Second, 2*time.Second, &styled.Theme{}, transport, peji.DefaultNotFound{})

	require.NoError(t, pages.Add("sales", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var sales = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		sales.AddLive("users", &UserComponent{})
		return sales
	}))

	require.NoError(t, pages.Add("about", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var about = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		about.AddLive("users", &UserComponent{})
		return about
	}))

	require.True(t, pages.Has("sales"))

	var salesRequest, salesReqErr = http.NewRequest("GET", "/sales", bytes.NewBuffer([]byte("alex")))
	require.NoError(t, salesReqErr)
	require.NotNil(t, salesRequest)

	var salesRequestMessage, salesRequestMessageErr = transposer.Transpose(salesRequest, sabuhp.Params{})
	require.NoError(t, salesRequestMessageErr)
	require.NotNil(t, salesRequestMessage)
	require.Equal(t, "/sales", salesRequestMessage.Path)

	var salesResponseErr = pages.Handle(salesRequestMessage, transport)
	require.NoError(t, salesResponseErr)

	var result = <-results
	require.NotNil(t, result)
	require.Equal(t, "text/html", result.ContentType)
	require.Equal(t, "/sales", result.Path)
	require.Len(t, result.Cookies, 1)
	require.Equal(t, result.Cookies[0].Name, peji.QueryAndCookieSessionIdName)
	require.NotEmpty(t, result.Headers.Get(peji.HeaderSessionIdName))
	require.Equal(t, result.Headers.Get(peji.HeaderSessionIdName), result.Cookies[0].Value)
	require.NotEmpty(t, result.Cookies[0].Value)

	var renderedOutput = string(result.Payload)
	require.NotEmpty(t, renderedOutput)
	require.NotContains(t, renderedOutput, "not found")
	require.Contains(t, renderedOutput, "user</div>")

	canceler()
	pages.Wait()
}

func TestPagesRouter_WithExistingSession_WithQuery(t *testing.T) {
	var newCtx, canceler = context.WithCancel(context.Background())

	var logger = new(LoggerImpl)
	var codec = &codecs.JsonCodec{}
	var transposer = sabuhp.NewCodecTransposer(codec, logger, -1)

	var results = make(chan *sabuhp.Message, 1)
	var listeners = map[string][]sabuhp.TransportResponse{}
	var transport = &testingutils.TransportImpl{
		ListenFunc: func(topic string, handler sabuhp.TransportResponse) sabuhp.Channel {
			listeners[topic] = append(listeners[topic], handler)
			return &testingutils.NoPubSubChannel{}
		},
		SendToAllFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			return nil
		},
		SendToOneFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			var targetListeners = listeners[data.Topic]
			if len(targetListeners) > 0 {
				_ = targetListeners[0].Handle(data, nil)
			}
			return nil
		},
	}

	var pages = peji.NewPages(newCtx, logger, "/", 5*time.Second, 2*time.Second, &styled.Theme{}, transport, peji.DefaultNotFound{})

	require.NoError(t, pages.Add("sales", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var sales = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		sales.AddLive("users", &UserComponent{})
		return sales
	}))

	require.NoError(t, pages.Add("about", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var about = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		about.AddLive("users", &UserComponent{})
		return about
	}))

	require.True(t, pages.Has("sales"))

	var salesRequest, salesReqErr = http.NewRequest("GET", "/sales", bytes.NewBuffer([]byte("alex")))
	require.NoError(t, salesReqErr)
	require.NotNil(t, salesRequest)

	var salesRequestMessage, salesRequestMessageErr = transposer.Transpose(salesRequest, sabuhp.Params{})
	require.NoError(t, salesRequestMessageErr)
	require.NotNil(t, salesRequestMessage)
	require.Equal(t, "/sales", salesRequestMessage.Path)

	var salesResponseErr = pages.Handle(salesRequestMessage, transport)
	require.NoError(t, salesResponseErr)

	var result = <-results
	require.NotNil(t, result)
	require.Equal(t, "text/html", result.ContentType)
	require.Equal(t, "/sales", result.Path)
	require.Len(t, result.Cookies, 1)
	require.Equal(t, result.Cookies[0].Name, peji.QueryAndCookieSessionIdName)
	require.NotEmpty(t, result.Headers.Get(peji.HeaderSessionIdName))
	require.Equal(t, result.Headers.Get(peji.HeaderSessionIdName), result.Cookies[0].Value)
	require.NotEmpty(t, result.Cookies[0].Value)

	var salesRequest2, salesReqErr2 = http.NewRequest(
		"GET",
		fmt.Sprintf(
			"/sales?%s=%s",
			peji.QueryAndCookieSessionIdName,
			result.Headers.Get(peji.HeaderSessionIdName),
		),
		bytes.NewBuffer([]byte("alex")),
	)
	require.NoError(t, salesReqErr2)
	require.NotNil(t, salesRequest2)

	var salesRequestMessage2, salesRequestMessageErr2 = transposer.Transpose(salesRequest2, sabuhp.Params{})
	require.NoError(t, salesRequestMessageErr2)
	require.NotNil(t, salesRequestMessage2)
	require.Equal(t, "/sales", salesRequestMessage2.Path)

	var salesResponseErr2 = pages.Handle(salesRequestMessage2, transport)
	require.NoError(t, salesResponseErr2)

	var result2 = <-results
	require.NotNil(t, result2)
	require.Equal(t, "text/html", result2.ContentType)
	require.Equal(t, "/sales", result2.Path)
	require.Len(t, result2.Cookies, 1)
	require.Equal(t, result2.Cookies[0].Name, peji.QueryAndCookieSessionIdName)
	require.NotEmpty(t, result2.Headers.Get(peji.HeaderSessionIdName))

	canceler()
	pages.Wait()
}

func TestPagesRouter_WithExistingSession_WithHeader(t *testing.T) {
	var newCtx, canceler = context.WithCancel(context.Background())

	var logger = new(LoggerImpl)
	var codec = &codecs.JsonCodec{}
	var transposer = sabuhp.NewCodecTransposer(codec, logger, -1)

	var results = make(chan *sabuhp.Message, 1)
	var listeners = map[string][]sabuhp.TransportResponse{}
	var transport = &testingutils.TransportImpl{
		ListenFunc: func(topic string, handler sabuhp.TransportResponse) sabuhp.Channel {
			listeners[topic] = append(listeners[topic], handler)
			return &testingutils.NoPubSubChannel{}
		},
		SendToAllFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			return nil
		},
		SendToOneFunc: func(data *sabuhp.Message, timeout time.Duration) error {
			results <- data
			var targetListeners = listeners[data.Topic]
			if len(targetListeners) > 0 {
				_ = targetListeners[0].Handle(data, nil)
			}
			return nil
		},
	}

	var pages = peji.NewPages(newCtx, logger, "/", 5*time.Second, 2*time.Second, &styled.Theme{}, transport, peji.DefaultNotFound{})

	require.NoError(t, pages.Add("sales", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var sales = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		sales.AddLive("users", &UserComponent{})
		return sales
	}))

	require.NoError(t, pages.Add("about", func(name string, theme *styled.Theme, pubsub sabuhp.Transport) *peji.Page {
		var about = peji.WithPage(name, theme, sampleLayout, peji.DefaultNotFound{})
		about.AddLive("users", &UserComponent{})
		return about
	}))

	require.True(t, pages.Has("sales"))

	var salesRequest, salesReqErr = http.NewRequest("GET", "/sales", bytes.NewBuffer([]byte("alex")))
	require.NoError(t, salesReqErr)
	require.NotNil(t, salesRequest)

	var salesRequestMessage, salesRequestMessageErr = transposer.Transpose(salesRequest, sabuhp.Params{})
	require.NoError(t, salesRequestMessageErr)
	require.NotNil(t, salesRequestMessage)
	require.Equal(t, "/sales", salesRequestMessage.Path)

	var salesResponseErr = pages.Handle(salesRequestMessage, transport)
	require.NoError(t, salesResponseErr)

	var result = <-results
	require.NotNil(t, result)
	require.Equal(t, "text/html", result.ContentType)
	require.Equal(t, "/sales", result.Path)
	require.Len(t, result.Cookies, 1)
	require.Equal(t, result.Cookies[0].Name, peji.QueryAndCookieSessionIdName)
	require.NotEmpty(t, result.Headers.Get(peji.HeaderSessionIdName))
	require.Equal(t, result.Headers.Get(peji.HeaderSessionIdName), result.Cookies[0].Value)
	require.NotEmpty(t, result.Cookies[0].Value)

	var salesRequest2, salesReqErr2 = http.NewRequest("GET", "/sales", bytes.NewBuffer([]byte("alex")))
	require.NoError(t, salesReqErr2)
	require.NotNil(t, salesRequest2)

	salesRequest2.Header.Set(peji.HeaderSessionIdName, result.Headers.Get(peji.HeaderSessionIdName))

	var salesRequestMessage2, salesRequestMessageErr2 = transposer.Transpose(salesRequest2, sabuhp.Params{})
	require.NoError(t, salesRequestMessageErr2)
	require.NotNil(t, salesRequestMessage2)
	require.Equal(t, "/sales", salesRequestMessage2.Path)
	require.Equal(t, result.Headers.Get(peji.HeaderSessionIdName), salesRequestMessage2.Headers.Get(peji.HeaderSessionIdName))

	var salesResponseErr2 = pages.Handle(salesRequestMessage2, transport)
	require.NoError(t, salesResponseErr2)

	var result2 = <-results
	require.NotNil(t, result2)
	require.Equal(t, "text/html", result2.ContentType)
	require.Equal(t, "/sales", result2.Path)
	require.Len(t, result2.Cookies, 1)
	require.Equal(t, result2.Cookies[0].Name, peji.QueryAndCookieSessionIdName)
	require.NotEmpty(t, result2.Headers.Get(peji.HeaderSessionIdName))

	canceler()
	pages.Wait()
}
