package main

import (
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

	r.GET("/", func(c *gin.Context) {
		err := RenderIndex("", c)
		if err != nil {
			panic(err)
		}
	})

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
