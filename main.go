package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"htmx/htmx"
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
		RenderToBody(c, RenderComponent("button", gin.H{}, Htmx{
			htmx.HxGet:     "/empty",
			htmx.HxSwap:    htmx.OuterHTML,
			htmx.HxPushUrl: "true",
		}, htmx.OnJs(htmx.Click, "./static/index.js")))
		return
	})

	r.GET("/empty", func(c *gin.Context) {
		RenderToBody(c, RenderComponent("empty", gin.H{}, Htmx{}))
	})

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
