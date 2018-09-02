package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

const REQUEST_ID = "X-Request-Id"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(REQUEST_ID)

		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Expose it for use teh application
		c.Set(REQUEST_ID, requestId)

		//Set X-Request-Id header
		c.Writer.Header().Set(REQUEST_ID, requestId)
		c.Next()
	}
}
