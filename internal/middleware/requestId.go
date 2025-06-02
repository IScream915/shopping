package middleware

import (
	"base_frame/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// gin.HandlerFunc 是Gin框架中用来表示中间件或路由处理函数的类型

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在ctx中设置一个唯一的operationID
		c.Set(constant.OperationID, uuid.NewString())
	}
}
