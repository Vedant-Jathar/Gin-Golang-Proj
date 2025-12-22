package middleware

import (
	"net/http"

	Utils "github.com/Vedant-Jathar/Gin_Project/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GlobalErrorHandlingMiddleware(logger zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			status := http.StatusInternalServerError
			message := "Internal server error"

			if appErr, ok := err.(*Utils.AppError); ok {
				status = appErr.StatusCode
				message = appErr.Message
			}

			logger.Error(message,
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", status),
			)

			c.JSON(status, gin.H{
				"status":  false,
				"message": message,
			})
		}
	}
}
