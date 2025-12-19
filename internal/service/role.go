package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type RoleService struct {
	pb.UnimplementedRoleServer
	uc *biz.RoleUsecase
}

func NewRoleService(uc *biz.RoleUsecase) *RoleService {
	return &RoleService{uc: uc}
}

func (s *RoleService) PageList(ctx context.Context, req *pb.PageRoleRequest) (*pb.PageRoleResponse, error) {
	total, roles, err := s.uc.PageList(ctx, int64(req.Page.PageIndex), int64(req.Page.PageSize))
	if err != nil {
		return nil, err
	}

	roleInfos := make([]*pb.RoleInfo, 0, len(roles))
	for _, role := range roles {
		roleInfos = append(roleInfos, &pb.RoleInfo{
			Id:        role.Id,
			Code:      role.Code,
			Name:      role.Name,
			CreatedAt: role.CreatedAt.Unix(),
			UpdatedAt: role.UpdatedAt.Unix(),
		})
	}

	return &pb.PageRoleResponse{
		Total: total,
		Roles: roleInfos,
	}, nil
}

func (s *RoleService) Get(ctx context.Context, req *pb.GetRoleRequest) (*pb.RoleInfo, error) {
	role, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.RoleInfo{
		Id:        role.Id,
		Code:      role.Code,
		Name:      role.Name,
		CreatedAt: role.CreatedAt.Unix(),
		UpdatedAt: role.UpdatedAt.Unix(),
	}, nil
}

func (s *RoleService) Create(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleInfo, error) {
	role := &biz.Role{
		Code: req.Code,
		Name: req.Name,
	}

	createdRole, err := s.uc.Create(ctx, role)
	if err != nil {
		return nil, err
	}

	return &pb.RoleInfo{
		Id:        createdRole.Id,
		Code:      createdRole.Code,
		Name:      createdRole.Name,
		CreatedAt: createdRole.CreatedAt.Unix(),
		UpdatedAt: createdRole.UpdatedAt.Unix(),
	}, nil
}

func (s *RoleService) Update(ctx context.Context, req *pb.UpdateRoleRequest) (*emptypb.Empty, error) {
	role := &biz.Role{
		ID:   biz.ID{Id: req.Id},
		Code: req.Code,
		Name: req.Name,
	}

	err := s.uc.Update(ctx, role)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *RoleService) Delete(ctx context.Context, req *pb.DeleteRoleRequest) (*emptypb.Empty, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
