package oppai

import (
	"net/http"
)

// HandlerFunc defines the request handler used by oppai
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of oppai.Engine
func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

// Add Routes
func (e *Engine) AddRoutes(method string, pattern string, handler HandlerFunc) {
	e.router.AddRoutes(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("POST", pattern, handler)
}

func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("PUT", pattern, handler)
}

func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("DELETE", pattern, handler)
}

func (e *Engine) OPTION(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("OPTION", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	e.router.handle(ctx)
}
