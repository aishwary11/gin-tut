package routes

import (
	"gin.com/aishwary11/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/login", controller.Login)
	}
}
