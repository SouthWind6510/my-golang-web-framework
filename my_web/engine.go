package my_web

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup
	router Router
	groups []*RouterGroup
}

func New(router Router) *Engine {
	if router == nil {
		router = NewDynamicRouter()
	}
	engine := &Engine{router: router}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := newContext(req, resp)
	e.router.handler(c)
}

func (e *Engine) Run(addr string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(addr, handler))
}
