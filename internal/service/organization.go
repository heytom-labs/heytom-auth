package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type OrganizationService struct {
	pb.UnimplementedOrganizationServer
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (s *OrganizationService) PageList(ctx context.Context, req *pb.PageOrganizationRequest) (*pb.PageOrganizationResponse, error) {
	return &pb.PageOrganizationResponse{}, nil
}
func (s *OrganizationService) Get(ctx context.Context, req *pb.GetOrganizationRequest) (*pb.OrganizationInfo, error) {
	return &pb.OrganizationInfo{}, nil
}
func (s *OrganizationService) Create(ctx context.Context, req *pb.CreateOrganizationRequest) (*pb.OrganizationInfo, error) {
	return &pb.OrganizationInfo{}, nil
}
func (s *OrganizationService) Update(ctx context.Context, req *pb.UpdateOrganizationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *OrganizationService) Delete(ctx context.Context, req *pb.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
