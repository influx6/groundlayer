package peji

import (
	"sync"
)

type PageNotification struct {
	hl       sync.RWMutex
	handlers []OnPage
}

func NewPageNotification() *PageNotification {
	return &PageNotification{}
}

func (h *PageNotification) Emit(route string, page *Page) {
	h.hl.RLock()
	defer h.hl.RUnlock()
	for _, handler := range h.handlers {
		handler(route, page)
	}
}

func (h *PageNotification) Add(eventHandler OnPage) {
	h.hl.Lock()
	h.handlers = append(h.handlers, eventHandler)
	h.hl.Unlock()
}
