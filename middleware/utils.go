package middleware

import (
	"github.com/gin-gonic/gin"
	"one-api/common"
	"one-api/common/logger"
)

func abortWithMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": gin.H{
			"message": common.MessageWithRequestId(message, c.GetString(common.RequestIdKey)),
			"type":    "one_api_error",
		},
	})
	c.Abort()
	logger.Error(c.Request.Context(), message)
}
