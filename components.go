package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"htmx/htmx"
)

type Htmx map[string]any

func RenderComponent(name string, params gin.H, htmx Htmx, callbacks ...*htmx.Callback) template.HTML {
	attr := make([]template.HTMLAttr, 0, len(htmx))
	for k, v := range htmx {
		attr = append(attr, template.HTMLAttr(fmt.Sprintf("%s=\"%s\"", k, v)))
	}

	for _, v := range callbacks {
		attr = append(attr, template.HTMLAttr(fmt.Sprintf("%s=\"%s\"", v.Event, v.Code)))
	}

	params["htmx"] = attr
	buff := &bytes.Buffer{}
	err := Templates.ExecuteTemplate(buff, name, params)
	if err != nil {
		fmt.Println(err)
		return template.HTML(fmt.Sprintf("Error while rendering %s", name))
	}
	return template.HTML(buff.String())
}

func RenderToBody(c *gin.Context, html template.HTML) {
	err := Templates.ExecuteTemplate(c.Writer, "index", gin.H{
		"Body": html,
	})
	if err != nil {
		panic(err)
	}
}

func Render(c *gin.Context, html template.HTML) error {
	_, err := io.WriteString(c.Writer, string(html))
	return err
}
