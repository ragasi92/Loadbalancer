package models

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy httputil.ReverseProxy
}

func (backend *Backend) SetAlive(alive bool) {
	backend.mux.Lock()
	backend.Alive = alive
	backend.mux.Unlock()
}

func (backend *Backend) IsAlive() (alive bool) {
	backend.mux.RLock()
	alive = backend.Alive
	backend.mux.RUnlock()
	return
}
