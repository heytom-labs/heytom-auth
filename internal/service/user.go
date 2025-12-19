package service

import (
	"context"
	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) PageList(ctx context.Context, req *pb.PageUserRequest) (*pb.PageUserResponse, error) {
	return &pb.PageUserResponse{}, nil
}
func (s *UserService) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
