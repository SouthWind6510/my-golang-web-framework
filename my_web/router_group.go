package my_web

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (rg *RouterGroup) Group(prefix string) *RouterGroup {
	engine := rg.engine
	newGroup := &RouterGroup{
		prefix: rg.prefix + prefix,
		parent: rg,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (rg *RouterGroup) addRouter(method string, path string, handler HandlerFunc) {
	rg.engine.router.addRouter(method, rg.prefix+path, handler)
}

func (rg *RouterGroup) GET(path string, handler HandlerFunc) {
	rg.addRouter("GET", path, handler)
}

func (rg *RouterGroup) POST(path string, handler HandlerFunc) {
	rg.addRouter("POST", path, handler)
}
