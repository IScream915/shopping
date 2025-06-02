package pcontext

import (
	"base_frame/internal/repo/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	CtxUserKey = "user" // 用户信息
)

// GetRequestToken 从请求头中获取token
func GetRequestToken(c *gin.Context) string {
	// 从请求头中获取Authorization参数
	authorization := c.Request.Header.Get("Authorization")
	// 如果Auth为空直接返回
	if authorization == "" {
		return ""
	}

	tokens := strings.SplitN(authorization, " ", 2)
	if len(tokens) == 2 || tokens[0] == "Bearer" {
		return tokens[1]
	}
	return ""
}

func GetUserTokenFromCtx(ctx context.Context) (*models.UserToken, error) {
	// 从context中拿到userToken
	user := ctx.Value(CtxUserKey)
	// 校验
	if user == nil {
		return nil, errors.New("user not found in context")
	}
	// token类型合法性
	userInfo, ok := user.(*models.UserToken)
	if !ok {
		return nil, errors.New("user not found in context")
	}
	// 无效的userID
	if userInfo.UserID == 0 {
		return nil, errors.New("user not found in context")
	}
	return userInfo, nil
}
