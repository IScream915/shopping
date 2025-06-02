package middleware

import (
	"base_frame/internal/repo"
	"base_frame/pkg/pcontext"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth(tokenRepo repo.UserToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从请求头中的Authorization获取Ticket
		ticket := c.GetHeader("Authorization")
		if ticket == "" {
			// 获取请求头中的Ticket失败后，尝试在从URL的参数中获取Ticket
			ticket, _ = c.GetQuery("token")
		}
		if ticket == "" {
			// TODO c.Abort, return, c.AbortWithStatusJSON 三者之间的关系和不同
			// 请求头和请求体中都没有token
			// c.JSON() 本身不会终止函数执行，它只是向 ResponseWriter 发送了 HTTP 响应内容
			// 但函数的执行仍然会继续，所以仍然需要 c.Abort() 和 return 来确保请求处理结束，不会执行后续的中间件或处理函数。
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			// c.Abort() 用于终止当前请求的后续处理，确保不执行后续的处理函数（handlers）或中间件
			// 如果没有 c.Abort()，Gin 仍然会继续执行后续的 c.Next() 处理流程
			c.Abort()
			// 这里可以优化成->
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			// 看起来更加的简洁
			return
		}

		// 根据ticket从redis中获取userInfo
		// 这里取出Authorization后，将Bearer与真正的Ticket分离，获取真正的Ticket
		userInfo, err := tokenRepo.Find(c, pcontext.GetRequestToken(c))
		if err != nil {
			log.Println("Authorization failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// 将从redis中获取的userInfo存入上下文中
		c.Set(pcontext.CtxUserKey, userInfo)
		// 处理下一个中间件，这里相当于是将当前的中间件压入栈中
		c.Next()
	}
}
