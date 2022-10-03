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
# Creating your context first
r := oppai.New()

# Add route to your root path
r.GET("/", func(ctx *oppai.Context) {
	ctx.Status(200)
})

# Start your server on port 3000
r.Run(":3000")
```
