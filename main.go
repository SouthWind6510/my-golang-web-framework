package main

import (
	"me/my-golang-web-framework/my_web"
	"me/my-golang-web-framework/my_web/router"
)

func main() {
	var r *router.Router
	r = router.NewStaticRouters()
	engine := my_web.New(r)
	register(engine)
	engine.Run(":9999", engine)
}
