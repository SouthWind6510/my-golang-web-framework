package my_web

type Router interface {
	addRouter(method, path string, handler HandlerFunc)
	handler(c *Context)
}
