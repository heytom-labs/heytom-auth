package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type OrganizationService struct {
	pb.UnimplementedOrganizationServer
	uc *biz.OrganizationUsecase
}

func NewOrganizationService(uc *biz.OrganizationUsecase) *OrganizationService {
	return &OrganizationService{uc: uc}
}

func (s *OrganizationService) PageList(ctx context.Context, req *pb.PageOrganizationRequest) (*pb.PageOrganizationResponse, error) {
	total, organizations, err := s.uc.PageList(ctx, int64(req.Page.PageIndex), int64(req.Page.PageSize))
	if err != nil {
		return nil, err
	}

	organizationInfos := make([]*pb.OrganizationInfo, 0, len(organizations))
	for _, organization := range organizations {
		organizationInfos = append(organizationInfos, &pb.OrganizationInfo{
			Id:        organization.Id,
			Name:      organization.Name,
			Code:      organization.Code,
			ParentId:  organization.ParentId,
			CreatedAt: organization.CreatedAt.Unix(),
			UpdatedAt: organization.UpdatedAt.Unix(),
		})
	}

	return &pb.PageOrganizationResponse{
		Total:         total,
		Organizations: organizationInfos,
	}, nil
}

func (s *OrganizationService) Get(ctx context.Context, req *pb.GetOrganizationRequest) (*pb.OrganizationInfo, error) {
	organization, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OrganizationInfo{
		Id:        organization.Id,
		Name:      organization.Name,
		Code:      organization.Code,
		ParentId:  organization.ParentId,
		CreatedAt: organization.CreatedAt.Unix(),
		UpdatedAt: organization.UpdatedAt.Unix(),
	}, nil
}

func (s *OrganizationService) Create(ctx context.Context, req *pb.CreateOrganizationRequest) (*pb.OrganizationInfo, error) {
	organization := &biz.Organization{
		Name:     req.Name,
		Code:     req.Code,
		ParentId: req.ParentId,
	}

	createdOrganization, err := s.uc.Create(ctx, organization)
	if err != nil {
		return nil, err
	}

	return &pb.OrganizationInfo{
		Id:        createdOrganization.Id,
		Name:      createdOrganization.Name,
		Code:      createdOrganization.Code,
		ParentId:  createdOrganization.ParentId,
		CreatedAt: createdOrganization.CreatedAt.Unix(),
		UpdatedAt: createdOrganization.UpdatedAt.Unix(),
	}, nil
}

func (s *OrganizationService) Update(ctx context.Context, req *pb.UpdateOrganizationRequest) (*emptypb.Empty, error) {
	organization := &biz.Organization{
		ID:       biz.ID{Id: req.Id},
		Name:     req.Name,
		Code:     req.Code,
		ParentId: req.ParentId,
	}

	err := s.uc.Update(ctx, organization)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrganizationService) Delete(ctx context.Context, req *pb.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
