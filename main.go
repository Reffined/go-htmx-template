package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var Templates *template.Template

func main() {
	glob, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	Templates = glob

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		RenderToBody(c, "index", gin.H{
			"Body": template.HTML(RenderComponent("button", gin.H{})),
		})
		return
	})
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
