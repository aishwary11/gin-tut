package routes

import (
	"gin.com/aishwary11/controller"
	"github.com/gin-gonic/gin"
)

func SetupItemRoutes(router *gin.Engine) {
	itemRoutes := router.Group("/item")
	{
		itemRoutes.GET("/", controller.GetItems)
		itemRoutes.GET("/:id", controller.GetItem)
	}
}
