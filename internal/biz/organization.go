package biz

import (
	"context"
)

type Organization struct {
	ID
	Name     string `json:"name" gorm:"size:30;not null;comment:名称"`
	Code     string `json:"code" gorm:"size:30;not null;comment:编号"`
	ParentId int64  `json:"parent_id"`
	Timestamps
	SoftDeletes
}

type OrganizationRepo interface {
	PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Organization, error)
	Get(ctx context.Context, id int64) (*Organization, error)
	Create(ctx context.Context, req *Organization) (*Organization, error)
	Update(ctx context.Context, req *Organization) error
	Delete(ctx context.Context, id int64) error
}

type OrganizationUsecase struct {
	repo OrganizationRepo
}

func NewOrganizationUsecase(repo OrganizationRepo) *OrganizationUsecase {
	return &OrganizationUsecase{repo: repo}
}

func (s *OrganizationUsecase) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*Organization, error) {
	return s.repo.PageList(ctx, pageIndex, pageSize)
}

func (s *OrganizationUsecase) Get(ctx context.Context, id int64) (*Organization, error) {
	return s.repo.Get(ctx, id)
}

func (s *OrganizationUsecase) Create(ctx context.Context, req *Organization) (*Organization, error) {
	return s.repo.Create(ctx, req)
}

func (s *OrganizationUsecase) Update(ctx context.Context, req *Organization) error {
	return s.repo.Update(ctx, req)
}

func (s *OrganizationUsecase) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
