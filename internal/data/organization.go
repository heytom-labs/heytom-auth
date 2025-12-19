package data

import (
	"context"

	"heytom-auth/internal/biz"
)

type Organization struct {
	ID
	Name     string `json:"name" gorm:"size:30;not null;comment:名称"`
	Code     string `json:"code" gorm:"size:30;not null;comment:编号"`
	ParentId int64  `json:"parent_id"`
	Timestamps
	SoftDeletes
}

type organizationRepo struct {
	data *Data
}

// NewOrganizationRepo .
func NewOrganizationRepo(data *Data) biz.OrganizationRepo {
	return &organizationRepo{
		data: data,
	}
}

func (r *organizationRepo) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*biz.Organization, error) {
	var organizations []*biz.Organization
	var total int64

	// 查询总数
	if err := r.data.db.Model(&Organization{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	// 分页查询
	offset := (pageIndex - 1) * pageSize
	if err := r.data.db.Offset(int(offset)).Limit(int(pageSize)).Find(&organizations).Error; err != nil {
		return 0, nil, err
	}

	return total, organizations, nil
}

func (r *organizationRepo) Get(ctx context.Context, id int64) (*biz.Organization, error) {
	var organization biz.Organization
	if err := r.data.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return nil, err
	}
	return &organization, nil
}

func (r *organizationRepo) Create(ctx context.Context, req *biz.Organization) (*biz.Organization, error) {
	if err := r.data.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *organizationRepo) Update(ctx context.Context, req *biz.Organization) error {
	if err := r.data.db.Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *organizationRepo) Delete(ctx context.Context, id int64) error {
	if err := r.data.db.Where("id = ?", id).Delete(&biz.Organization{}).Error; err != nil {
		return err
	}
	return nil
}
