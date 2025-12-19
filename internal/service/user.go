package service

import (
	"context"
	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) PageList(ctx context.Context, req *pb.PageUserRequest) (*pb.PageUserResponse, error) {
	total, users, err := s.uc.PageList(ctx, int64(req.Page.PageIndex), int64(req.Page.PageSize))
	if err != nil {
		return nil, err
	}
	
	userInfos := make([]*pb.UserInfo, 0, len(users))
	for _, user := range users {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:        user.Id,
			Name:      user.Name,
			Mobile:    user.Mobile,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		})
	}
	
	return &pb.PageUserResponse{
		Total: total,
		Users: userInfos,
	}, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	user, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	
	return &pb.UserInfo{
		Id:        user.Id,
		Name:      user.Name,
		Mobile:    user.Mobile,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}, nil
}

func (s *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	user := &biz.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}
	
	result, err := s.uc.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return &pb.UserInfo{
		Id:        result.Id,
		Name:      result.Name,
		Mobile:    result.Mobile,
		CreatedAt: result.CreatedAt.Unix(),
		UpdatedAt: result.UpdatedAt.Unix(),
	}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {
	user := &biz.User{
		Id:     req.Id,
		Name:   req.Name,
		Mobile: req.Mobile,
	}
	
	err := s.uc.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	
	return &emptypb.Empty{}, nil
}
