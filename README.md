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
r := oppai.New()

// Add route to your root path for GET request and return with HTTP status 200 OK
r.GET("/", func(ctx *oppai.Context) {
	ctx.Status(200)
})

// Start your server and listen to 0.0.0.0:3000, you could change this with your desired address
r.Run(":3000")

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

### Oppai Route
```go
func (e *RouterGroup) GET(pattern string, handler HandlerFunc)
```

Serve Oppai with GET Method on certain route defined in pattern.

### net/http

For more information, please read on https://pkg.go.dev/net/http#pkg-constants
