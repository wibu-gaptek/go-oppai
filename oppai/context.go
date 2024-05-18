package oppai

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

	// Middlewares
	handlers []HandlerFunc
	index    int

	// config
	oppaiCfg *OppaiConfig
}

func newContext(w http.ResponseWriter, r *http.Request, config *OppaiConfig) *Context {
	return &Context{
		Writer:   w,
		Req:      r,
		Path:     r.URL.Path,
		Method:   r.Method,
		index:    -1,
		oppaiCfg: config,
	}
}

func (ctx *Context) Next() {
	ctx.index++
	skip := len(ctx.handlers)

	for ; ctx.index < skip; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

func (ctx *Context) Fail(code int, err string) {
	ctx.index = len(ctx.handlers)
	ctx.JSON(code, H{
		"error": err,
	})
}

func (ctx *Context) Param(key string) string {
	return ctx.Params[key]
}

func (ctx *Context) setgeneralHeader() {
	if ctx.oppaiCfg.GeneralHeader {
		ctx.Writer.Header().Set("Powered-By", "Go Oppai Engine")
	}
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
	ctx.setgeneralHeader()
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, object interface{}) {
	ctx.setgeneralHeader()
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)

	encode := json.NewEncoder(ctx.Writer)
	if err := encode.Encode(object); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.setgeneralHeader()
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.setgeneralHeader()
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}

func (ctx *Context) Bind(d any) error {
	ctx.SetHeader("Content-Type", "application/json")
	decoder := json.NewDecoder(ctx.Req.Body)
	return decoder.Decode(d)
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.Writer, cookie)
}

func (ctx *Context) GetCookie(name string) string {
	cookie, err := ctx.Req.Cookie(name)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(ctx.Writer, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(ctx.Writer, "server error", http.StatusInternalServerError)
		}
	}

	return cookie.Value
}
