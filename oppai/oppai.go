package oppai

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by oppai
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}

	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// New is the constructor of oppai.Engine
func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

// Group is defined to create a new RouterGroup
// remember all groups share the same Engine instance
func (e *RouterGroup) Group(prefix string) *RouterGroup {
	engine := e.engine
	newGroup := &RouterGroup{
		prefix: e.prefix + prefix,
		parent: e,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

// Use is defined to add middleware to the group
func (e *RouterGroup) Use(middlewares ...HandlerFunc) {
	e.middlewares = append(e.middlewares, middlewares...)
}

func (e *RouterGroup) AddRoutes(method string, comp string, handler HandlerFunc) {
	pattern := e.prefix + comp
	e.engine.router.AddRoutes(method, pattern, handler)
}

func (e *RouterGroup) GET(pattern string, handler HandlerFunc) {
	e.AddRoutes("GET", pattern, handler)
}

func (e *RouterGroup) POST(pattern string, handler HandlerFunc) {
	e.AddRoutes("POST", pattern, handler)
}

func (e *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	e.AddRoutes("PUT", pattern, handler)
}

func (e *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	e.AddRoutes("DELETE", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}
