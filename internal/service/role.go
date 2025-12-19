package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type RoleService struct {
	pb.UnimplementedRoleServer
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) PageList(ctx context.Context, req *pb.PageRoleRequest) (*pb.PageRoleResponse, error) {
	return &pb.PageRoleResponse{}, nil
}
func (s *RoleService) Get(ctx context.Context, req *pb.GetRoleRequest) (*pb.RoleInfo, error) {
	return &pb.RoleInfo{}, nil
}
func (s *RoleService) Create(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleInfo, error) {
	return &pb.RoleInfo{}, nil
}
func (s *RoleService) Update(ctx context.Context, req *pb.UpdateRoleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *RoleService) Delete(ctx context.Context, req *pb.DeleteRoleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
