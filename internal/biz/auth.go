package biz

import (
	"context"
	"time"
)

// Token 令牌信息
type Token struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	UserId       int64     `json:"user_id" gorm:"not null;index;comment:用户ID"`
	Token        string    `json:"token" gorm:"size:500;not null;comment:访问令牌"`
	RefreshToken string    `json:"refresh_token" gorm:"size:500;not null;comment:刷新令牌"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null;comment:过期时间"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// AuthRepo 认证仓储接口
type AuthRepo interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *User) (*User, error)
	// GetUserByMobile 根据手机号获取用户
	GetUserByMobile(ctx context.Context, mobile string) (*User, error)
	// GetUserById 根据ID获取用户
	GetUserById(ctx context.Context, id int64) (*User, error)
	// CreateToken 创建令牌
	CreateToken(ctx context.Context, token *Token) (*Token, error)
	// GetTokenByUserId 根据用户ID获取令牌
	GetTokenByUserId(ctx context.Context, userId int64) (*Token, error)
	// DeleteToken 删除令牌
	DeleteToken(ctx context.Context, userId int64) error
}

// AuthUsecase 认证用例
type AuthUsecase struct {
	repo AuthRepo
}

// NewAuthUsecase 新建认证用例
func NewAuthUsecase(repo AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

// Register 注册
func (au *AuthUsecase) Register(ctx context.Context, name, mobile, password string) (*Token, error) {
	// 先检查用户是否已存在
	existingUser, err := au.repo.GetUserByMobile(ctx, mobile)
	if err == nil && existingUser != nil {
		// 用户已存在，返回错误或相应处理
		// 这里简化处理，实际应该返回具体的业务错误
		return nil, nil
	}

	// 创建新用户
	user := &User{
		Name:     name,
		Mobile:   mobile,
		Password: password, // 实际应该加密存储
	}
	createdUser, err := au.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// 创建令牌
	token := &Token{
		UserId:       createdUser.Id,
		Token:        "access_token_placeholder", // 实际应该生成真实的JWT令牌
		RefreshToken: "refresh_token_placeholder",
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7), // 7天过期
	}
	createdToken, err := au.repo.CreateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return createdToken, nil
}

// Login 登录
func (au *AuthUsecase) Login(ctx context.Context, mobile, password string) (*Token, error) {
	// 获取用户
	user, err := au.repo.GetUserByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}

	// 验证密码（实际应该使用加密验证）
	if user.Password != password {
		return nil, nil // 密码错误
	}

	// 检查是否已有有效令牌
	existingToken, err := au.repo.GetTokenByUserId(ctx, user.Id)
	if err == nil && existingToken != nil && existingToken.ExpiresAt.After(time.Now()) {
		// 返回现有有效令牌
		return existingToken, nil
	}

	// 创建新令牌
	token := &Token{
		UserId:       user.Id,
		Token:        "access_token_placeholder", // 实际应该生成真实的JWT令牌
		RefreshToken: "refresh_token_placeholder",
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7), // 7天过期
	}
	createdToken, err := au.repo.CreateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return createdToken, nil
}

// Logout 登出
func (au *AuthUsecase) Logout(ctx context.Context, userId int64) error {
	// 删除用户的令牌
	return au.repo.DeleteToken(ctx, userId)
}

// LoginByGithub 通过GitHub登录
func (au *AuthUsecase) LoginByGithub(ctx context.Context, code string) (*Token, error) {
	// 这里应该实现GitHub OAuth流程
	// 1. 使用code换取access_token
	// 2. 使用access_token获取用户信息
	// 3. 查找或创建用户
	// 4. 生成并返回令牌

	// 简化实现，实际应替换为真实的GitHub OAuth逻辑
	user := &User{
		Name:         "github_user",
		Mobile:       "",    // GitHub登录可能没有手机号
		GithubOpenid: 12345, // 示例GitHub ID
		Avatar:       "https://github.com/images/avatar.jpg",
	}

	// 查找或创建用户
	// 这里简化处理，实际应该有更复杂的逻辑
	createdUser, err := au.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// 创建令牌
	token := &Token{
		UserId:       createdUser.Id,
		Token:        "access_token_placeholder", // 实际应该生成真实的JWT令牌
		RefreshToken: "refresh_token_placeholder",
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7), // 7天过期
	}
	createdToken, err := au.repo.CreateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return createdToken, nil
}
