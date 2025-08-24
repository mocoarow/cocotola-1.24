package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewWaitMiddleware() gin.HandlerFunc {
	return func(_ *gin.Context) {
		time.Sleep(time.Second)
	}
}
