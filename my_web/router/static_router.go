package router

import (
	"log"
	"me/my-golang-web-framework/my_web"
	"net/http"
	"reflect"
)

var r *Router = new(Router)

type staticRouter struct {
	router map[string]my_web.HandlerFunc
}

func NewStaticRouters() *staticRouter {
	return &staticRouter{
		router: make(map[string]my_web.HandlerFunc),
	}
}

func (r *staticRouter) addRouter(method, path string, handler my_web.HandlerFunc) {
	key := method + "_" + path
	// 同一个路径绑定不同的handler方法
	if _, ok := r.router[key]; ok && !reflect.DeepEqual(r.router[key], handler) {
		log.Printf("%v already exist", key)
		return
	}
	r.router[key] = handler
}

func (r *staticRouter) handler(c *my_web.Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.router[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "%s not found", key)
	}
}
