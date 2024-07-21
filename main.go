package main

import (
	"fmt"
	"htmx/htmx"
	"text/template"

	"github.com/gin-gonic/gin"
)

var Templates *template.Template

func main() {
	Templates = template.New("button")
	_, err := Templates.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET(
		"/", func(c *gin.Context) {
			body := RenderComp("login", Params{
				"MakeHtmx": MakeHtmx,
				"htmx": Htmx{
					htmx.HxPost:   "/empty",
					htmx.HxTarget: "body",
				},
				"id": "login",
				"submit": Params{
					"MakeHtmx": MakeHtmx,
					"htmx":     Htmx{},
					"formId":   "login",
					"class":    "flg1",
				},
			})

			err := RenderIndex(body, c)
			if err != nil {
				panic(err)
			}
		},
	)

	r.POST(
		"/empty", func(c *gin.Context) {
			email, ok := c.GetPostForm("email")
			if ok {
				fmt.Println(email)
			}

			err := Render(c, RenderComp("empty", Params{"txt": email}))
			if err != nil {
				panic(err)
			}
		},
	)

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
