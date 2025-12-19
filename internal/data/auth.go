package data

import (
	"context"
	"heytom-auth/internal/biz"
)

type authRepo struct {
	data *Data
}

// NewAuthRepo .
func NewAuthRepo(data *Data) biz.AuthRepo {
	return &authRepo{
		data: data,
	}
}

// CreateUser 创建用户
func (r *authRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	if err := r.data.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByMobile 根据手机号获取用户
func (r *authRepo) GetUserByMobile(ctx context.Context, mobile string) (*biz.User, error) {
	var user biz.User
	if err := r.data.db.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserById 根据ID获取用户
func (r *authRepo) GetUserById(ctx context.Context, id int64) (*biz.User, error) {
	var user biz.User
	if err := r.data.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateToken 创建令牌
func (r *authRepo) CreateToken(ctx context.Context, token *biz.Token) (*biz.Token, error) {
	if err := r.data.db.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

// GetTokenByUserId 根据用户ID获取令牌
func (r *authRepo) GetTokenByUserId(ctx context.Context, userId int64) (*biz.Token, error) {
	var token biz.Token
	if err := r.data.db.Where("user_id = ?", userId).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

// DeleteToken 删除令牌
func (r *authRepo) DeleteToken(ctx context.Context, userId int64) error {
	if err := r.data.db.Where("user_id = ?", userId).Delete(&biz.Token{}).Error; err != nil {
		return err
	}
	return nil
}
