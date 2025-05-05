package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func GetHistoryHandler(c *gin.Context) {
	ctx := context.Background()
	ip := c.ClientIP()
	history := service.GetHistory(ctx, ip)
	c.JSON(http.StatusOK, history)
}
