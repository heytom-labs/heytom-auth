package data

import (
	"context"

	"heytom-auth/internal/biz"
)

type Role struct {
	ID
	Code string `json:"code" gorm:"size:30;not null;comment:编号"`
	Name string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Timestamps
	SoftDeletes
}

type roleRepo struct {
	data *Data
}

// NewRoleRepo .
func NewRoleRepo(data *Data) biz.RoleRepo {
	return &roleRepo{
		data: data,
	}
}

func (r *roleRepo) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*biz.Role, error) {
	var roles []*biz.Role
	var total int64

	// 查询总数
	if err := r.data.db.Model(&Role{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	// 分页查询
	offset := (pageIndex - 1) * pageSize
	if err := r.data.db.Offset(int(offset)).Limit(int(pageSize)).Find(&roles).Error; err != nil {
		return 0, nil, err
	}

	return total, roles, nil
}

func (r *roleRepo) Get(ctx context.Context, id int64) (*biz.Role, error) {
	var role biz.Role
	if err := r.data.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) Create(ctx context.Context, req *biz.Role) (*biz.Role, error) {
	if err := r.data.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *roleRepo) Update(ctx context.Context, req *biz.Role) error {
	if err := r.data.db.Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	if err := r.data.db.Where("id = ?", id).Delete(&biz.Role{}).Error; err != nil {
		return err
	}
	return nil
}
