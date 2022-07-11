package main

import (
	"me/my-golang-web-framework/handler"
	"me/my-golang-web-framework/my_web"
)

func register(engine *my_web.Engine) {
	engine.GET("/ping", handler.Ping)
	engine.POST("/login", handler.Login)
	engine.GET("/hello", handler.Hello)
}
