package handler

import (
	"fmt"
	"me/my-golang-web-framework/my_web"
	"net/http"
)

func Ping(c *my_web.Context) {
	c.String(http.StatusOK, "hello %s\n", c.Query("username"))
}

func Login(c *my_web.Context) {
	c.JSON(http.StatusOK, my_web.JsonBody{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

func Hello(c *my_web.Context) {
	html := fmt.Sprintf("<h1>hello %s</h1>", c.Query("username"))
	c.HTML(http.StatusOK, html)
}

func Name(c *my_web.Context) {
	c.String(http.StatusOK, "hello %s\nthere are %s\n", c.Params["name"], c.Path)
}
