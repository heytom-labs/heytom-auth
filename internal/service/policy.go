package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type PolicyService struct {
	pb.UnimplementedPolicyServer
	uc *biz.PolicyUsecase
}

func NewPolicyService(uc *biz.PolicyUsecase) *PolicyService {
	return &PolicyService{uc: uc}
}

func (s *PolicyService) PageList(ctx context.Context, req *pb.PagePolicyRequest) (*pb.PagePolicyResponse, error) {
	total, policies, err := s.uc.PageList(ctx, int64(req.Page.PageIndex), int64(req.Page.PageSize))
	if err != nil {
		return nil, err
	}

	policyInfos := make([]*pb.PolicyInfo, 0, len(policies))
	for _, policy := range policies {
		policyInfos = append(policyInfos, &pb.PolicyInfo{
			Id:          policy.Id,
			Name:        policy.Name,
			Description: policy.Description,
			CreatedAt:   policy.CreatedAt.Unix(),
			UpdatedAt:   policy.UpdatedAt.Unix(),
		})
	}

	return &pb.PagePolicyResponse{
		Total:     total,
		Policies:  policyInfos,
	}, nil
}

func (s *PolicyService) Get(ctx context.Context, req *pb.GetPolicyRequest) (*pb.PolicyInfo, error) {
	policy, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.PolicyInfo{
		Id:          policy.Id,
		Name:        policy.Name,
		Description: policy.Description,
		CreatedAt:   policy.CreatedAt.Unix(),
		UpdatedAt:   policy.UpdatedAt.Unix(),
	}, nil
}

func (s *PolicyService) Create(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.PolicyInfo, error) {
	policy := &biz.Policy{
		Name:        req.Name,
		Description: req.Description,
	}

	createdPolicy, err := s.uc.Create(ctx, policy)
	if err != nil {
		return nil, err
	}

	return &pb.PolicyInfo{
		Id:          createdPolicy.Id,
		Name:        createdPolicy.Name,
		Description: createdPolicy.Description,
		CreatedAt:   createdPolicy.CreatedAt.Unix(),
		UpdatedAt:   createdPolicy.UpdatedAt.Unix(),
	}, nil
}

func (s *PolicyService) Update(ctx context.Context, req *pb.UpdatePolicyRequest) (*emptypb.Empty, error) {
	policy := &biz.Policy{
		ID:          biz.ID{Id: req.Id},
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.uc.Update(ctx, policy)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *PolicyService) Delete(ctx context.Context, req *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
