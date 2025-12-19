package biz

import (
	"context"
)

type Application struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	AppKey      string `json:"app_key" gorm:"size:50;not null;comment:appkey"`
	AppSecret   string `json:"app_secret" gorm:"size:100;not null;comment:密钥"`
	CallbackUrl string `json:"callback_url" gorm:"size:255;not null;comment:登录回调地址"`
	Timestamps
	SoftDeletes
}

type ApplicationRepo interface {
	PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Application, error)
	Get(ctx context.Context, id int64) (*Application, error)
	Create(ctx context.Context, req *Application) (*Application, error)
	Update(ctx context.Context, req *Application) error
	Delete(ctx context.Context, id int64) error
}

type ApplicationUsecase struct {
	repo ApplicationRepo
}

func NewApplicationUsecase(repo ApplicationRepo) *ApplicationUsecase {
	return &ApplicationUsecase{repo: repo}
}

func (s *ApplicationUsecase) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Application, error) {
	return s.repo.PageList(ctx, pageIndex, pageSize)
}

func (s *ApplicationUsecase) Get(ctx context.Context, id int64) (*Application, error) {
	return s.repo.Get(ctx, id)
}

func (s *ApplicationUsecase) Create(ctx context.Context, req *Application) (*Application, error) {
	return s.repo.Create(ctx, req)
}

func (s *ApplicationUsecase) Update(ctx context.Context, req *Application) error {
	return s.repo.Update(ctx, req)
}

func (s *ApplicationUsecase) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
