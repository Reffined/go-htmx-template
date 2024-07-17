package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RenderComponent(name string, params gin.H) string {
	buff := &bytes.Buffer{}
	err := Templates.ExecuteTemplate(buff, name, params)
	if err != nil {
		return fmt.Sprintf("Error while rendering %s", name)
	}
	return buff.String()
}

func RenderToBody(c *gin.Context, name string, params gin.H) {
	err := Templates.ExecuteTemplate(c.Writer, name, params)
	if err != nil {
		panic(err)
	}
}
