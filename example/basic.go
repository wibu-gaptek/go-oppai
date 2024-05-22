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

type todoPageData struct {
	PageTitle string
	Todos     []todos
}

type todos struct {
	Title string
	Done  bool
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

	r.GET("/tmpl", func(ctx *oppai.Context) {
		data := todoPageData{
			PageTitle: "My TODO list",
			Todos: []todos{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		if err := ctx.RenderHTML("example/assets/index.html", data); err != nil {
			ctx.JSON(http.StatusBadRequest, oppai.H{
				"err": err,
			})
		}
	})

	r.GET("/ip", func(ctx *oppai.Context) {
		ip := ctx.GetIPAdress()
		ctx.JSON(http.StatusBadRequest, oppai.H{
			"ip": ip,
		})
	})

	r.Run(":9000")
}
