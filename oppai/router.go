package oppai

import (
	"net/http"
)

type router struct {
	root     map[string]*node
	handlers map[string]HandlerFunc
}

type oppaiHandlerCfg struct {
	OppaiCtx *Context
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		root:     make(map[string]*node),
	}
}

func (r *router) AddRoutes(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "_" + pattern

	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{}
	}

	r.root[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
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
			} else if part[0] == '*' {
				params[part[1:]] = joinParts(searchParts[i:])
				break
			}
		}
		return n, params
	}
	return nil, nil
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

func joinParts(parts []string) string {
	length := 0
	for _, part := range parts {
		length += len(part) + 1
	}
	joined := make([]byte, length-1)
	pos := 0
	for _, part := range parts {
		copy(joined[pos:], part)
		pos += len(part)
		if pos < length-1 {
			joined[pos] = '/'
			pos++
		}
	}
	return string(joined)
}
