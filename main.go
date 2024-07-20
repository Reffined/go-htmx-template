package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
	"htmx/errors"
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
			err := RenderIndex(RenderComp("button", nil, nil), c)
			if err != nil {
				errors.InternalServerError(c, err)
				return
			}
			return
		},
	)

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
