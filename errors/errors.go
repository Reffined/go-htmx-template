package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
