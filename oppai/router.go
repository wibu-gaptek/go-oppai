package oppai

import (
	"net/http"
	"strings"
)

type router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

type oppaiHandlerCfg struct {
	// ctx
	OppaiCtx *Context
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		root:     make(map[string]*node),
	}
}

func (r *router) AddRoutes(method string, pattern string, handler HandlerFunc) {

	parts := parsePattern(pattern)

	key := method + "_" + pattern
	_, ok := r.root[method]

	if !ok {
		r.root[method] = &node{}
	}

	r.root[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.root[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}

			if part[0] == '*' && len(part) > 0 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}

		return n, params
	}

	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.root[method]

	if !ok {
		return nil
	}

	nodes := make([]*node, 0)
	root.travel(&nodes)

	return nodes
}

func (r *router) handle(ctx *oppaiHandlerCfg) {
	n, params := r.getRoute(ctx.OppaiCtx.Method, ctx.OppaiCtx.Path)
	if n != nil {
		ctx.OppaiCtx.Params = params
		key := ctx.OppaiCtx.Method + "_" + n.pattern
		r.handlers[key](ctx.OppaiCtx)
	} else {
		http.Error(ctx.OppaiCtx.Writer, "NOT FOUND", http.StatusNotFound)
	}
}

