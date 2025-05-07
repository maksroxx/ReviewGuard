package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func GetHistoryHandler(svc service.HistoryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		history := svc.GetHistory(ctx, ip)
		c.JSON(http.StatusOK, history)
	}
}
