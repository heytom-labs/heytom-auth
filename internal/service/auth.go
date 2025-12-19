package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TokenResponse, error) {
	return &pb.TokenResponse{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	return &pb.TokenResponse{}, nil
}
func (s *AuthService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *AuthService) LoginByGithub(ctx context.Context, req *pb.LoginByGithubRequest) (*pb.TokenResponse, error) {
	return &pb.TokenResponse{}, nil
}
