package services

import (
	"base_frame/global"
	"base_frame/internal/dto"
	"base_frame/internal/repo"
	"base_frame/internal/repo/models"
	"base_frame/pkg/email"
	"base_frame/pkg/pcontext"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User interface {
	AccountLogin(ctx context.Context, req *dto.AccountLoginReq) (*models.UserToken, error)
	EmailSend(ctx context.Context, req *dto.EmailSendReq) error
	EmailLogin(ctx context.Context, req *dto.EmailLoginReq) (*models.UserToken, error)
	Logout(ctx context.Context, ticket string) error
	Create(ctx context.Context, req *dto.CreateUserReq) error
	Update(ctx context.Context, req *dto.UpdateUserReq) error
	Delete(ctx context.Context, req *dto.DeleteUserReq) error
}

func NewUser(repo repo.User, tokenRepo repo.UserToken, gCfg *global.Config) User {
	return &user{
		tokenRepo: tokenRepo,
		repo:      repo,
		gCfg:      gCfg,
	}
}

type user struct {
	tokenRepo repo.UserToken
	repo      repo.User
	gCfg      *global.Config
}

func (obj *user) AccountLogin(ctx context.Context, req *dto.AccountLoginReq) (*models.UserToken, error) {
	// 根据账号查找用户信息
	user, err := obj.repo.FindByAccount(ctx, req.Account)
	if err != nil {
		return nil, errors.New("账号错误，请输入正确的账号")
	}
	// 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误，请重新输入密码")
	}
	fmt.Println("salt: ", obj.gCfg.Salt.Secret)
	// 保存用户登录信息到redis
	tokenInfo := models.UserToken{
		UserID:    user.ID,
		Account:   user.Account,
		Nickname:  user.NickName,
		ExpiredAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		Ticket:    uuid.NewString(),
	}
	err = obj.tokenRepo.Save(ctx, &tokenInfo)
	if err != nil {
		return nil, errors.New("登录凭证保存失败")
	}
	return &tokenInfo, nil
}

func (obj *user) EmailSend(ctx context.Context, req *dto.EmailSendReq) error {
	// 生成验证码并发送
	code, err := email.SendCodeEmail(req.Email)
	if err != nil {
		return errors.New("验证码发送失败")
	}
	tokenInfo := models.UserToken{
		VerificationCode: code,
		ExpiredAt:        time.Now().Add(time.Minute * 5).Unix(),
		Ticket:           req.Email,
	}
	// 将验证码保存到redis中，便于后续的验证
	err = obj.tokenRepo.Save(ctx, &tokenInfo)
	if err != nil {
		return errors.New("邮箱验证码保存失败")
	}
	return nil
}

func (obj *user) EmailLogin(ctx context.Context, req *dto.EmailLoginReq) (*models.UserToken, error) {
	// 从redis中获取邮箱登录凭证
	userToken, err := obj.tokenRepo.Find(ctx, req.Email)
	if err != nil {
		return nil, errors.New("邮箱错误，无法获取登录信息，请输入正确的邮箱")
	}
	// 校验验证码
	if userToken.VerificationCode != req.VerificationCode {
		return nil, errors.New("验证码错误，请输入正确的验证码")
	}
	// 验证通过，根据邮箱获取用户详细信息
	user, err := obj.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("获取用户信息失败，请重新检查邮箱")
	}
	// 删除redis中的验证码记录
	err = obj.tokenRepo.Delete(ctx, req.Email)
	if err != nil {
		return nil, errors.New("删除验证码凭证失败")
	}
	// 保存用户登录信息到redis
	tokenInfo := models.UserToken{
		UserID:           user.ID,
		Account:          user.Account,
		Nickname:         user.NickName,
		VerificationCode: req.VerificationCode,
		ExpiredAt:        time.Now().Add(time.Hour * 24 * 7).Unix(),
		Ticket:           uuid.NewString(),
	}
	err = obj.tokenRepo.Save(ctx, &tokenInfo)
	if err != nil {
		return nil, errors.New("登录凭证保存失败")
	}
	return &tokenInfo, nil
}

func (obj *user) Logout(ctx context.Context, ticket string) error {
	if err := obj.tokenRepo.Delete(ctx, ticket); err != nil {
		return errors.New("登出失败")
	}
	return nil
}

func (obj *user) Create(ctx context.Context, req *dto.CreateUserReq) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("生成密码哈希失败")
	}
	user := models.User{
		Account:  req.Account,
		Email:    req.Email,
		NickName: req.NickName,
		Password: string(hashPassword),
		Age:      req.Age,
		Sex:      req.Sex,
	}
	txErr := obj.repo.Transaction(ctx, func(tx *gorm.DB) error {
		err := obj.repo.WithTx(tx).Create(ctx, &user)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}
	return nil
}

func (obj *user) Update(ctx context.Context, req *dto.UpdateUserReq) error {
	userToken, err := pcontext.GetUserTokenFromCtx(ctx)
	if err != nil {
		return err
	}
	newUser := &models.User{}
	newUser.ID = userToken.UserID
	newUser.Email = req.Email
	newUser.NickName = req.NickName
	newUser.Password = req.Password
	newUser.Age = req.Age
	newUser.Sex = req.Sex

	txErr := obj.repo.Transaction(ctx, func(tx *gorm.DB) error {
		err = obj.repo.WithTx(tx).Update(ctx, newUser)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	return nil
}

func (obj *user) Delete(ctx context.Context, req *dto.DeleteUserReq) error {
	txErr := obj.repo.Transaction(ctx, func(tx *gorm.DB) error {
		err := obj.repo.DeleteByIds(ctx, req.IDs)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}
	return nil
}
