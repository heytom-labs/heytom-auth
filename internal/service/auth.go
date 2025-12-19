package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	au *biz.AuthUsecase
}

func NewAuthService(au *biz.AuthUsecase) *AuthService {
	return &AuthService{
		au: au,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.TokenResponse, error) {
	token, err := s.au.Register(ctx, req.Name, req.Mobile, req.Password)
	if err != nil || token == nil {
		return nil, err
	}

	return &pb.TokenResponse{
		AccessToken:  token.Token,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresAt.Unix(),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	token, err := s.au.Login(ctx, req.Mobile, req.Password)
	if err != nil || token == nil {
		return nil, err
	}

	return &pb.TokenResponse{
		AccessToken:  token.Token,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresAt.Unix(),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	// 在实际实现中，我们需要从上下文中获取用户ID
	// 这里简化处理，假设用户ID为1
	err := s.au.Logout(ctx, 1)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *AuthService) LoginByGithub(ctx context.Context, req *pb.LoginByGithubRequest) (*pb.TokenResponse, error) {
	token, err := s.au.LoginByGithub(ctx, req.Code)
	if err != nil || token == nil {
		return nil, err
	}

	return &pb.TokenResponse{
		AccessToken:  token.Token,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresAt.Unix(),
	}, nil
}
