package main

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
)

type Htmx map[string]any

func RenderIndex(body string, c *gin.Context) error {
	err := Templates.ExecuteTemplate(
		c.Writer, "index", gin.H{
			"body": body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func RenderComp(name string, params gin.H, htmx Htmx) string {
	buf := bytes.Buffer{}
	if params != nil && htmx != nil {
		params["htmx"] = htmx
	}
	err := Templates.ExecuteTemplate(&buf, name, params)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func Render(c *gin.Context, res string) error {
	_, err := io.WriteString(c.Writer, res)
	if err != nil {
		return err
	}
	return nil
}
