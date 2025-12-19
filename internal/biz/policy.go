package biz

import (
	"context"
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

type PolicyRepo interface {
	PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Policy, error)
	Get(ctx context.Context, id int64) (*Policy, error)
	Create(ctx context.Context, req *Policy) (*Policy, error)
	Update(ctx context.Context, req *Policy) error
	Delete(ctx context.Context, id int64) error
}

type PolicyUsecase struct {
	repo PolicyRepo
}

func NewPolicyUsecase(repo PolicyRepo) *PolicyUsecase {
	return &PolicyUsecase{repo: repo}
}

func (s *PolicyUsecase) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Policy, error) {
	return s.repo.PageList(ctx, pageIndex, pageSize)
}

func (s *PolicyUsecase) Get(ctx context.Context, id int64) (*Policy, error) {
	return s.repo.Get(ctx, id)
}

func (s *PolicyUsecase) Create(ctx context.Context, req *Policy) (*Policy, error) {
	return s.repo.Create(ctx, req)
}

func (s *PolicyUsecase) Update(ctx context.Context, req *Policy) error {
	return s.repo.Update(ctx, req)
}

func (s *PolicyUsecase) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
