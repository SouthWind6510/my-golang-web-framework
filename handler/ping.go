package handler

import (
	"fmt"
	"net/http"
)

func Ping(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "pong")
}

func Pong(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "pong")
}
