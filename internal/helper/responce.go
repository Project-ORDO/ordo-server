package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Responce(c *gin.Context, status int, message string, data interface{}, redirect string) {
	c.JSON(status, gin.H{
		"status":   http.StatusText(status),
		"message":  message,
		"code":     status,
		"data":     data,
		"redirect": redirect,
	})
}
