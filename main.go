package main

import (
	"log"
	"os"

	"gin.com/aishwary11/middleware"
	"gin.com/aishwary11/routes"
	"gin.com/aishwary11/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.ConnectDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RateLimitMiddleware())
	routes.SetupUserRoutes(router)
	router.Use(middleware.JwtAuthMiddleware())
	routes.SetupItemRoutes(router)
	router.Run(":" + port)
}
