package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        for _, err := range c.Errors {
            logger.Error(err)
        }
    }
}