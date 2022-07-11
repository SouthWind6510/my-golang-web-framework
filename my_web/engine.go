package my_web

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routerMap map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{routerMap: make(map[string]HandlerFunc)}
}

func (e *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := req.Method + "_" + req.URL.Path
	if handler, ok := e.routerMap[key]; ok {
		handler(resp, req)
	} else {
		log.Printf("%v not find", key)
		resp.WriteHeader(404)
		fmt.Fprintf(resp, "%v not found\n", key)
	}
}

func (e *Engine) addRouter(method, path string, handler HandlerFunc) {
	key := method + "_" + path
	// 同一个路径绑定不同的handler方法
	if _, ok := e.routerMap[key]; ok && !reflect.DeepEqual(e.routerMap[key], handler) {
		log.Printf("%v already exist", key)
		return
	}
	e.routerMap[key] = handler
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.addRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandlerFunc) {
	e.addRouter("POST", path, handler)
}

func (e *Engine) Run(addr string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(addr, handler))
}
