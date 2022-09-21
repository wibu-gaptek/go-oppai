package oppai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {

	// Origin Objects from Http Lib
	Writer http.ResponseWriter
	Req    *http.Request

	// Request Info
	Path   string
	Method string
	Params map[string]string

	// Response Info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]

	return value
}

func (ctx *Context) GeneralHeader() {
	ctx.Writer.Header().Set("Powered-By", "Go Oppai Engine")
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.GeneralHeader()
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, object interface{}) {
	ctx.GeneralHeader()
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)

	encode := json.NewEncoder(ctx.Writer)
	if err := encode.Encode(object); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.GeneralHeader()
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.GeneralHeader()
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
