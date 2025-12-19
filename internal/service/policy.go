package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type PolicyService struct {
	pb.UnimplementedPolicyServer
}

func NewPolicyService() *PolicyService {
	return &PolicyService{}
}

func (s *PolicyService) PageList(ctx context.Context, req *pb.PagePolicyRequest) (*pb.PagePolicyResponse, error) {
	return &pb.PagePolicyResponse{}, nil
}
func (s *PolicyService) Get(ctx context.Context, req *pb.GetPolicyRequest) (*pb.PolicyInfo, error) {
	return &pb.PolicyInfo{}, nil
}
func (s *PolicyService) Create(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.PolicyInfo, error) {
	return &pb.PolicyInfo{}, nil
}
func (s *PolicyService) Update(ctx context.Context, req *pb.UpdatePolicyRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PolicyService) Delete(ctx context.Context, req *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
