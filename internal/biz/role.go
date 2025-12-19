package biz

import (
	"context"
)

type Role struct {
	ID
	Code string `json:"code" gorm:"size:30;not null;comment:编号"`
	Name string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Timestamps
	SoftDeletes
}

type RoleRepo interface {
	PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Role, error)
	Get(ctx context.Context, id int64) (*Role, error)
	Create(ctx context.Context, req *Role) (*Role, error)
	Update(ctx context.Context, req *Role) error
	Delete(ctx context.Context, id int64) error
}

type RoleUsecase struct {
	repo RoleRepo
}

func NewRoleUsecase(repo RoleRepo) *RoleUsecase {
	return &RoleUsecase{repo: repo}
}

func (s *RoleUsecase) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Role, error) {
	return s.repo.PageList(ctx, pageIndex, pageSize)
}

func (s *RoleUsecase) Get(ctx context.Context, id int64) (*Role, error) {
	return s.repo.Get(ctx, id)
}

func (s *RoleUsecase) Create(ctx context.Context, req *Role) (*Role, error) {
	return s.repo.Create(ctx, req)
}

func (s *RoleUsecase) Update(ctx context.Context, req *Role) error {
	return s.repo.Update(ctx, req)
}

func (s *RoleUsecase) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
