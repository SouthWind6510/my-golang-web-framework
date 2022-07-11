package my_web

import (
	"log"
	"net/http"
	"reflect"
)

type HandlerFunc func(*Context)

type routers struct {
	router map[string]HandlerFunc
}

func newRouters() *routers {
	return &routers{
		router: make(map[string]HandlerFunc),
	}
}

func (r *routers) addRouter(method, path string, handler HandlerFunc) {
	key := method + "_" + path
	// 同一个路径绑定不同的handler方法
	if _, ok := r.router[key]; ok && !reflect.DeepEqual(r.router[key], handler) {
		log.Printf("%v already exist", key)
		return
	}
	r.router[key] = handler
}

func (r *routers) handler(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.router[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "%s not found", key)
	}
}
