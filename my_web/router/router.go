package router

import "me/my-golang-web-framework/my_web"

type Router interface {
	addRouter(method, path string, handler my_web.HandlerFunc)
	handler(c *my_web.Context)
}
