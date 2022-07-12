package my_web

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router Router
}

func New(router Router) *Engine {
	return &Engine{router: router}
}

func (e *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := newContext(req, resp)
	e.router.handler(c)
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.router.addRouter("POST", path, handler)
}

func (e *Engine) Run(addr string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(addr, handler))
}
