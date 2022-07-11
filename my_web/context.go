package my_web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter

	Path   string
	Method string

	StatusCode int
}

func newContext(req *http.Request, resp http.ResponseWriter) *Context {
	return &Context{
		Req:    req,
		Resp:   resp,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Resp.Write([]byte(fmt.Sprintf(format, values...)))
}

type JsonBody map[string]interface{}

func (c *Context) JSON(code int, obj JsonBody) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Resp, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Resp.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Resp.Write(data)
}
