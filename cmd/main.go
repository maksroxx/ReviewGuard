package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/handlers"
	redisclient "github.com/maksroxx/ReviewGuard/internal/redis"
)

func main() {
	redisclient.InitRedis()
	router := gin.Default()

	router.POST("/review", handlers.ReviewHandler)
	router.GET("/moderation/history", handlers.GetHistoryHandler)

	router.Run(":8080")
}
