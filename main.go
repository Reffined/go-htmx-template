package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
	"htmx/errors"
	"htmx/htmx"
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
			err := RenderIndex(
				RenderComp(
					"double", Params{
						"bot1": Params{
							"MakeHtmx": MakeHtmx,
							"htmx": Htmx{
								htmx.HxGet:     "/empty",
								htmx.HxTrigger: "click",
								htmx.HxTarget:  "body",
							},
							"text": "Submit",
						},
						"bot2": Params{
							"MakeHtmx": MakeHtmx,
							"htmx": Htmx{
								htmx.HxGet:     "/empty",
								htmx.HxTrigger: "click",
								htmx.HxTarget:  "body",
							},
							"text": "Test",
						},
					},
				), c,
			)
			if err != nil {
				errors.InternalServerError(c, err)
				return
			}
			return
		},
	)

	r.GET(
		"/empty", func(c *gin.Context) {
			err := Render(c, RenderComp("empty", nil))
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
