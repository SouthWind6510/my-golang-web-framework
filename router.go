package main

import (
	"me/my-golang-web-framework/handler"
	"me/my-golang-web-framework/my_web"
)

func register(engine *my_web.Engine) {
	engine.GET("/ping", handler.Ping)
	engine.GET("/ping", handler.Pong)
}
