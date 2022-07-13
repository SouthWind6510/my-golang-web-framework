package main

import (
	"me/my-golang-web-framework/main/handler"
	"me/my-golang-web-framework/my_web"
)

func register(engine *my_web.Engine) {
	engine.GET("/ping", handler.Ping)
	engine.POST("/login", handler.Login)
	engine.GET("/hello", handler.Hello)
	engine.GET("/name/:name", handler.Name)
}
