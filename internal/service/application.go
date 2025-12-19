package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ApplicationService struct {
	pb.UnimplementedApplicationServer
}

func NewApplicationService() *ApplicationService {
	return &ApplicationService{}
}

func (s *ApplicationService) PageList(ctx context.Context, req *pb.PageApplicationRequest) (*pb.PageApplicationResponse, error) {
	return &pb.PageApplicationResponse{}, nil
}
func (s *ApplicationService) Get(ctx context.Context, req *pb.GetApplicationRequest) (*pb.ApplicationInfo, error) {
	return &pb.ApplicationInfo{}, nil
}
func (s *ApplicationService) Create(ctx context.Context, req *pb.CreateApplicationRequest) (*pb.ApplicationInfo, error) {
	return &pb.ApplicationInfo{}, nil
}
func (s *ApplicationService) Update(ctx context.Context, req *pb.UpdateApplicationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *ApplicationService) Delete(ctx context.Context, req *pb.DeleteApplicationRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
