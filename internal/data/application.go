package data

import (
	"context"

	"heytom-auth/internal/biz"
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

type applicationRepo struct {
	data *Data
}

// NewApplicationRepo .
func NewApplicationRepo(data *Data) biz.ApplicationRepo {
	return &applicationRepo{
		data: data,
	}
}

func (r *applicationRepo) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*biz.Application, error) {
	var applications []*biz.Application
	var total int64

	// 查询总数
	if err := r.data.db.Model(&Application{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	// 分页查询
	offset := (pageIndex - 1) * pageSize
	if err := r.data.db.Offset(int(offset)).Limit(int(pageSize)).Find(&applications).Error; err != nil {
		return 0, nil, err
	}

	return total, applications, nil
}

func (r *applicationRepo) Get(ctx context.Context, id int64) (*biz.Application, error) {
	var application biz.Application
	if err := r.data.db.Where("id = ?", id).First(&application).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *applicationRepo) Create(ctx context.Context, req *biz.Application) (*biz.Application, error) {
	if err := r.data.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *applicationRepo) Update(ctx context.Context, req *biz.Application) error {
	if err := r.data.db.Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *applicationRepo) Delete(ctx context.Context, id int64) error {
	if err := r.data.db.Where("id = ?", id).Delete(&biz.Application{}).Error; err != nil {
		return err
	}
	return nil
}
