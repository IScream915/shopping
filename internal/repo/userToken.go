package repo

import (
	"base_frame/internal/errcode"
	"base_frame/internal/repo/models"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// UserToken中保存了用户的基础公开信息，并且设置了一个ticket值
// 在redis中将使用UserToken中的ticket值作为 key 用于进行索引
// 而 key 对应 的payload则是序列化之后的UserToken

type UserToken interface {
	// Save 将用户信息保存进redis
	Save(ctx context.Context, info *models.UserToken) error
	// Find 在redis中根据票据查找用户信息
	Find(ctx context.Context, ticket string) (*models.UserToken, error)
	// Delete 在redis中删除用户信息
	Delete(ctx context.Context, ticket string) error
}

func NewUserToken(client redis.UniversalClient) UserToken {
	return &userToken{client: client}
}

type userToken struct {
	client redis.UniversalClient
}

func (obj *userToken) Save(ctx context.Context, info *models.UserToken) error {
	// 将用户信息序列化
	data, err := info.MarshalBinary()
	if err != nil {
		return err
	}

	// 判断token的到期时间是否错误，如果到期时间在当前时间之前，那么到期时间设置错误
	now := time.Now().Unix()
	if info.ExpiredAt <= now {
		return errcode.EntityParameterError
	}

	// 存入redis
	// 这里time.Duration的基础单位是纳秒，而time.Second表示1秒对应的纳秒数
	return obj.client.Set(ctx, info.Ticket, data, time.Duration(info.ExpiredAt-now)*time.Second).Err()
}

func (obj *userToken) Find(ctx context.Context, ticket string) (*models.UserToken, error) {
	// 根据ticket从redis中取出用户信息
	var info models.UserToken
	// Scan 方法用于将 Redis 返回的数据反序列化为 Go 语言中的数据结构
	// 这里的 Scan 方法可以使用redis-tag来指定反序列化方向
	err := obj.client.Get(ctx, ticket).Scan(&info)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errcode.DBRecordNotFound
		}
		return nil, err
	}
	return &info, nil
}

func (obj *userToken) Delete(ctx context.Context, ticket string) error {
	// 根据ticket删除redis中指定的userToken内容
	return obj.client.Del(ctx, ticket).Err()
}
