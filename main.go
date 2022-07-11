package main

import (
	"me/my-golang-web-framework/my_web"
)

func main() {
	engine := my_web.New()
	register(engine)
	engine.Run(":9999", engine)
}
