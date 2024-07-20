package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type Htmx map[string]any
type Params map[string]any

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

func RenderComp(name string, params Params) string {
	buf := bytes.Buffer{}
	err := Templates.ExecuteTemplate(&buf, name, params)
	if err != nil {
		fmt.Println(err)
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

func MakeHtmx(h Htmx) string {
	buf := strings.Builder{}
	for k, v := range h {
		buf.WriteString(fmt.Sprintf(" %s=\"%s\"", k, v))
	}
	return buf.String()
}
