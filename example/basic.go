package main

import (
	"fmt"
	"net/http"

	"github.com/wibu-gaptek/go-oppai/oppai"
)

type userRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := oppai.New(oppai.OppaiConfig{
		GeneralHeader: true,
		DebugMode:     true,
	})

	r.GET("/say/:name", func(ctx *oppai.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, oppai.H{
			"name": name,
		})
	})

	r.POST("/register", func(ctx *oppai.Context) {
		user := userRegister{}
		if err := ctx.Bind(&user); err != nil {
			fmt.Printf("user: %v\n", user)
			return
		}
		fmt.Printf("user: %v\n", user)
		ctx.JSON(http.StatusOK, oppai.H{
			"email":    user.Email,
			"password": user.Password,
		})
	})

	r.GET("/set", func(ctx *oppai.Context) {
		cookie := new(http.Cookie)
		cookie.Name = "test-cookie"
		cookie.Value = "rhovhnnvinvirovrlnvknlnbkfnknknsknkfdfdjfosjfo"
		cookie.Path = "/"
		cookie.MaxAge = 3600
		cookie.HttpOnly = true
		cookie.Secure = true
		cookie.SameSite = http.SameSiteLaxMode
		ctx.SetCookie(cookie)

		ctx.JSON(200, oppai.H{
			"message": "success set cookie",
		})
	})

	r.GET("/get", func(ctx *oppai.Context) {
		cookieValue := ctx.GetCookie("test-cookie")

		if cookieValue != nil {
			ctx.JSON(200, oppai.H{
				"message": "success get cookie",
				"cookie":  cookieValue,
			})
		}

	})

	r.Run(":9000")
}
