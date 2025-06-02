package middleware

import (
	"base_frame/pkg/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 分配一个operationID
		c.Set(constant.OperationID, uuid.NewString())

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		//log.Info(c, fmt.Sprintf("%s: %s %v", c.Request.Method, path, cost),
		//	"status", c.Writer.Status(),
		//	"method", c.Request.Method,
		//	"path", path,
		//	"query", query,
		//	"ip", c.ClientIP(),
		//	"user-agent", c.Request.UserAgent(),
		//	"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		//	"cost", cost,
		//)
		log.Println(fmt.Sprintf("%s: %s %v", c.Request.Method, path, query))
		log.Println("status:", c.Writer.Status())
		log.Println("method:", c.Request.Method)
		log.Println("path:", path)
		log.Println("query:", query)
		log.Println("ip:", c.ClientIP())
		log.Println("user-agent:", c.Request.UserAgent())
		log.Println("error:", c.Errors.ByType(gin.ErrorTypePrivate).String())
		log.Println("cost:", cost)
	}
}
