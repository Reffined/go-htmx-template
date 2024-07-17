package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

func RenderComponent(name string, params gin.H, htmx gin.H) template.HTML {
	attr := make([]template.HTMLAttr, 0, len(htmx))
	for k, v := range htmx {
		attr = append(attr, template.HTMLAttr(fmt.Sprintf("%s=\"%s\"", k, v)))
	}
	params["htmx"] = attr
	buff := &bytes.Buffer{}
	err := Templates.ExecuteTemplate(buff, name, params)
	if err != nil {
		return template.HTML(fmt.Sprintf("Error while rendering %s", name))
	}
	return template.HTML(buff.String())
}

func RenderToBody(c *gin.Context, name string, params gin.H) {
	err := Templates.ExecuteTemplate(c.Writer, name, params)
	if err != nil {
		panic(err)
	}
}
