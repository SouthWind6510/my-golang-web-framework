package main

import (
	"me/my-golang-web-framework/main/handler"
	"me/my-golang-web-framework/my_web"
)

func register(engine *my_web.Engine) {
	engine.GET("/ping", handler.Ping)
	v1 := engine.Group("/v1")
	v1.GET("/hello", handler.Hello)
	api := v1.Group("/api")
	api.POST("/login", handler.Login)
	api.GET("/name/:name", handler.Name)
}
