package biz

import (
	"context"
	"time"
)

type User struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Mobile       string    `json:"mobile" gorm:"size:24;not null;index;comment:用户手机号"`
	Password     string    `json:"-" gorm:"not null;default:'';comment:用户密码"`
	GithubOpenid int64     `json:"github_openid" gorm:"comment:github openid"`
	Avatar       string    `json:"avatar" gorm:"size:255;not null;comment:头像"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepo interface {
	PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*User, error)
	Get(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, req *User) (*User, error)
	Update(ctx context.Context, req *User) error
	Delete(ctx context.Context, id int64) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (s *UserUsecase) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*User, error) {
	return s.repo.PageList(ctx, pageIndex, pageSize)
}
func (s *UserUsecase) Get(ctx context.Context, id int64) (*User, error) {
	return s.repo.Get(ctx, id)
}
func (s *UserUsecase) Create(ctx context.Context, req *User) (*User, error) {
	return s.repo.Create(ctx, req)
}
func (s *UserUsecase) Update(ctx context.Context, req *User) error {
	return s.repo.Update(ctx, req)
}
func (s *UserUsecase) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
