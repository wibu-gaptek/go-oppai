package oppai

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var oppaiLogo = ` _____             _ 
|     |___ ___ ___|_|
|  |  | . | . | .'| |
|_____|  _|  _|__,|_|
      |_| |_|                   
`

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
		config OppaiConfig
	}

	OppaiConfig struct {
		GeneralHeader bool
		DebugMode     bool
	}
)

// New is the constructor of oppai.Engine
func New(cfg ...OppaiConfig) *Engine {
	var opaiCfg OppaiConfig
	if len(cfg) > 0 {
		opaiCfg = cfg[0]
	}

	engine := &Engine{router: NewRouter(), config: opaiCfg}
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

func (e *Engine) OPTION(pattern string, handler HandlerFunc) {
	e.router.AddRoutes("OPTION", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	println(oppaiLogo)
	if e.engine.config.DebugMode {
		fmt.Println("DEBUGMODE IS ACTIVE")
	}

	if isPortInUse(addr) {
		panic("port was used!")
	}

	log.Printf("http server running on %s", addr)
	return http.ListenAndServe(addr, e)

}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req, &e.config)
	ctx.handlers = middlewares
	oppaiHandlerCfg := &oppaiHandlerCfg{ctx}
	e.router.handle(oppaiHandlerCfg)
}

func isPortInUse(port string) bool {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return true // return true if port was use
	}
	ln.Close()
	return false
}
