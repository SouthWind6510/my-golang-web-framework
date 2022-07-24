package main

import "me/my-golang-web-framework/my_web"

func main() {
	engine := my_web.New(nil)
	register(engine)
	engine.Run(":9999", engine)
}
