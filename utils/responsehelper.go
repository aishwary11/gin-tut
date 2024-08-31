package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHelper(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status < http.StatusBadRequest,
		"message": message,
		"data":    data,
	})
}
