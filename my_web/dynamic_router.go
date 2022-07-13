package my_web

import (
	"net/http"
	"strings"
)

var _ Router = new(dynamicRouter)

type dynamicRouter struct {
	router map[string]HandlerFunc
	roots  map[string]*node
}

func NewDynamicRouter() *dynamicRouter {
	return &dynamicRouter{
		router: make(map[string]HandlerFunc),
		roots:  make(map[string]*node),
	}
}

func parsePath(path string) (result []string) {
	parts := strings.Split(path, "/")
	result = make([]string, 0)
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return
}

func (r dynamicRouter) addRouter(method, path string, handler HandlerFunc) {
	_, ok := r.roots[method]
	// 请求方法的根结点
	if !ok {
		r.roots[method] = &node{}
	}
	// 向 trie 树中添加路由
	parts := parsePath(path)
	r.roots[method].insert(path, parts, 0)
	r.router[method+"_"+path] = handler
}

func (r dynamicRouter) getRouter(c *Context) (*node, map[string]string) {
	method := c.Method
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	parts := parsePath(c.Path)
	node := root.search(parts, 0)
	if node == nil {
		return nil, nil
	}
	parts2 := parsePath(node.path)
	param := make(map[string]string, 0)
	for index, part := range parts2 {
		if part[0] == ':' {
			param[part[1:]] = parts[index]
		}
	}
	return node, param
}

func (r dynamicRouter) handler(c *Context) {
	node, param := r.getRouter(c)
	if node != nil {
		c.Params = param
		key := c.Method + "_" + node.path
		r.router[key](c)
	} else {
		key := c.Method + "_" + c.Path
		c.String(http.StatusNotFound, "%s not found", key)
	}
}
