package main

import "me/my-golang-web-framework/my_web"

func main() {
	//r := my_web.NewStaticRouters()
	r := my_web.NewDynamicRouter()
	engine := my_web.New(r)
	register(engine)
	engine.Run(":9999", engine)
}
