package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/db"
	"github.com/maksroxx/ReviewGuard/internal/handlers"
	redisclient "github.com/maksroxx/ReviewGuard/internal/redis"
	"github.com/maksroxx/ReviewGuard/internal/service"
	"github.com/maksroxx/ReviewGuard/internal/worker"
)

func main() {
	var (
		redis    = redisclient.NewRedisClient()
		postgres = db.NewPostgreDB("postgres://user:password@localhost:5432/reviewguard?sslmode=disable")
		router   = gin.Default()

		historySvc = service.NewHistoryService(*redis)
		cacheSvc   = service.NewCacheService(*redis)
		rep        = db.NewPostgresRepository(*postgres)
	)

	router.POST("/review", handlers.ReviewHandler(*redis, *historySvc, *rep))
	router.GET("/moderation/history", handlers.GetHistoryHandler(*historySvc))
	router.GET("/report/stats", handlers.ReportStatsHandler(*cacheSvc, *rep))
	router.GET("/report/spam", handlers.GetSpamReviewsHandler(*rep))
	router.GET("/report/by-ip", handlers.GetReviewsByIPHandler(*rep))

	go worker.StartModerationWorker(redis.RDB, rep)

	router.Run(":8080")
}
