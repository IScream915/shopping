package api

import (
	"base_frame/internal/dto"
	"base_frame/internal/services"
	"base_frame/pkg/pcontext"
	"base_frame/pkg/response"
	"github.com/gin-gonic/gin"
)

type User interface {
	AccountLogin(c *gin.Context)
	EmailSend(c *gin.Context)
	EmailLogin(c *gin.Context)
	Logout(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUser(svc services.User) User {
	return &user{svc: svc}
}

type user struct {
	svc services.User
}

// AccountLogin 采用账号登录
func (obj *user) AccountLogin(c *gin.Context) {
	req := &dto.AccountLoginReq{}
	if err := c.ShouldBind(req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	userToken, err := obj.svc.AccountLogin(c, req)
	if err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithData(userToken))
}

// EmailSend 通过邮箱发送验证码
func (obj *user) EmailSend(c *gin.Context) {
	req := &dto.EmailSendReq{}
	if err := c.ShouldBind(req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	err := obj.svc.EmailSend(c, req)
	if err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

// EmailLogin 采用邮箱验证码形式登录
func (obj *user) EmailLogin(c *gin.Context) {
	req := &dto.EmailLoginReq{}
	if err := c.ShouldBind(req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	userToken, err := obj.svc.EmailLogin(c, req)
	if err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithData(userToken))
}

// Logout 账号登出
func (obj *user) Logout(c *gin.Context) {
	if err := obj.svc.Logout(c, pcontext.GetRequestToken(c)); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

// Create 注册用户
func (obj *user) Create(c *gin.Context) {
	req := &dto.CreateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Create(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

// Update 用户信息更新
func (obj *user) Update(c *gin.Context) {
	req := &dto.UpdateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Update(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

func (obj *user) Delete(c *gin.Context) {
	req := &dto.DeleteUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Delete(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}
	response.Json(c, response.WithMsg("success"))
}
