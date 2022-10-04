# Go Oppai

Lightweight Web Framework, written in Golang Native HTTP.

![OPPAI](https://steamuserimages-a.akamaihd.net/ugc/847091252323981832/2D70011394A10CE3D03D1347E1C298FA8F12FBFA/?imw=5000&imh=5000&ima=fit&impolicy=Letterbox&imcolor=#000000&letterbox=false)

## Install

```go get -u github.com/wibu-gaptek/go-oppai/oppai```

## Get Started

Start by importing go-oppai to your project
```go
import (
	"net/http"
	"github.com/wibu-gaptek/go-oppai/oppai"
)
```

Server example running on port 3000
```go
// Construct oppai first
oppaiServer := oppai.New()

// Add route to your root path for GET request and return with HTTP status 200 OK
oppaiServer.GET("/", func(ctx *oppai.Context) {
	ctx.Status(http.StatusOk)
})

// Start your server and listen to 0.0.0.0:3000, you could change this with your desired address
oppaiServer.Run(":3000")

// or if you want listen to 192.168.0.1:8080
r.Run("192.168.0.1:8080")
```

# Docs
### Oppai Constructor
```go
func New() *Engine
```
Construct Oppai engine, it automatically initiate router engine.

### Oppai Run
```go
func (e *Engine) Run(addr string) error
```

Run listens on the TCP network address **addr** and then calls Serve to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

If **addr** is blank, ":http" (:80) is used. 

### Oppai Router
```go
// GET Method
func (e *RouterGroup) GET(pattern string, handler HandlerFunc)

// POST Method
func (e *RouterGroup) POST(pattern string, handler HandlerFunc)

// PUT Method
func (e *RouterGroup) PUT(pattern string, handler HandlerFunc)

// DELETE Method
func (e *RouterGroup) DELETE(pattern string, handler HandlerFunc)

// OPTION Method
func (e *RouterGroup) OPTION(pattern string, handler HandlerFunc)

// by default, go-oppai had built with GET, POST, PUT, DELETE, OPTION
// However, you could add more route with another defined method
func (e *RouterGroup) AddRoutes(method string, comp string, handler HandlerFunc)
```

Serve Oppai with specific Method on certain route defined in pattern.

### Oppai Route group
```go
func (e *RouterGroup) Group(prefix string) *RouterGroup
```

Serve Oppai Router 

### net/http

For more information, please read on https://pkg.go.dev/net/http#pkg-constants

## Examples

### Add parameter to router
```go
// Construct oppai first
oppaiServer := oppai.New()

// Catch any request under /echo like /echo/hello or /echo/world et cetera
oppaiServer.GET("/echo/:name", func(ctx *oppai.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, oppai.H{
		"name": name,
	})
})

// Listen to  port 3000
oppaiServer.Run(":3000")
```

### Router Group
With router group, you could subroute like  normal
```go

// Construct oppai first
oppaiServer := oppai.New()

group := oppaiServer.Group("/echo")
{
	group.GET("/", func(ctx *oppai.Context) {
		ctx.JSON(http.StatusOK, oppai.H{
			"message": "Subroute with Router Group",
		})
	})
}

// Listen to  port 3000
oppaiServer.Run(":3000")
```
