package data

import (
	"context"

	"heytom-auth/internal/biz"
)

type Policy struct {
	ID
	Name        string `json:"name" gorm:"size:20;not null;comment:名称"`
	Description string `json:"description" gorm:"size:100;not null;comment:描述"`
	Timestamps
	SoftDeletes
}

type PolicyAuth struct {
	ID
	PolicyId   int64 `json:"policy_id"`
	TargetType int8  `json:"target_type" gorm:"comment:0用户1角色2部门"`
	TargetId   int64 `json:"target_id"`
}

type PolicyResource struct {
	ID
	PolicyId          int64 `json:"policy_id"`
	PermissionSpaceId int64 `json:"permission_space_id"`
	ResourceId        int64 `json:"resource_id"`
	Effect            int8  `json:"effect" gorm:"comment:效果0/1"`
	Timestamps
	SoftDeletes
}

type PolicyResourceItem struct {
	ID
	PolicyId            int64  `json:"policy_id"`
	PolicyResourceId    int64  `json:"policy_resouce_id"`
	ResourceItemId      int64  `json:"resource_item_id"`
	ResourceItemActions string `json:"resource_item_actions" gorm:"size:255;not null"`
}

type policyRepo struct {
	data *Data
}

// NewPolicyRepo .
func NewPolicyRepo(data *Data) biz.PolicyRepo {
	return &policyRepo{
		data: data,
	}
}

func (r *policyRepo) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*biz.Policy, error) {
	var policies []*biz.Policy
	var total int64

	// 查询总数
	if err := r.data.db.Model(&Policy{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	// 分页查询
	offset := (pageIndex - 1) * pageSize
	if err := r.data.db.Offset(int(offset)).Limit(int(pageSize)).Find(&policies).Error; err != nil {
		return 0, nil, err
	}

	return total, policies, nil
}

func (r *policyRepo) Get(ctx context.Context, id int64) (*biz.Policy, error) {
	var policy biz.Policy
	if err := r.data.db.Where("id = ?", id).First(&policy).Error; err != nil {
		return nil, err
	}
	return &policy, nil
}

func (r *policyRepo) Create(ctx context.Context, req *biz.Policy) (*biz.Policy, error) {
	if err := r.data.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *policyRepo) Update(ctx context.Context, req *biz.Policy) error {
	if err := r.data.db.Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *policyRepo) Delete(ctx context.Context, id int64) error {
	if err := r.data.db.Where("id = ?", id).Delete(&biz.Policy{}).Error; err != nil {
		return err
	}
	return nil
}
