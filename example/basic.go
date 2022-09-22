package main

import (
	"net/http"

	"github.com/wibu-gaptek/go-oppai/oppai"
)

func main() {
	r := oppai.New()

	r.GET("/", func(ctx *oppai.Context) {
		ctx.Status(200)
	})

	r.GET("/say/:name", func(ctx *oppai.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, oppai.H{
			"name": name,
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *oppai.Context) {
			ctx.JSON(http.StatusOK, oppai.H{
				"message": "Api V1",
			})
		})
	}

	r.Run(":3000")
}
