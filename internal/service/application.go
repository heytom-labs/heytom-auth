package service

import (
	"context"

	pb "heytom-auth/api/auth/v1"
	"heytom-auth/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ApplicationService struct {
	pb.UnimplementedApplicationServer
	uc *biz.ApplicationUsecase
}

func NewApplicationService(uc *biz.ApplicationUsecase) *ApplicationService {
	return &ApplicationService{uc: uc}
}

func (s *ApplicationService) PageList(ctx context.Context, req *pb.PageApplicationRequest) (*pb.PageApplicationResponse, error) {
	total, applications, err := s.uc.PageList(ctx, int64(req.Page.PageIndex), int64(req.Page.PageSize))
	if err != nil {
		return nil, err
	}

	applicationInfos := make([]*pb.ApplicationInfo, 0, len(applications))
	for _, application := range applications {
		applicationInfos = append(applicationInfos, &pb.ApplicationInfo{
			Id:          application.Id,
			Name:        application.Name,
			Description: application.Description,
			AppKey:      application.AppKey,
			AppSecret:   application.AppSecret,
			CallbackUrl: application.CallbackUrl,
			CreatedAt:   application.CreatedAt.Unix(),
			UpdatedAt:   application.UpdatedAt.Unix(),
		})
	}

	return &pb.PageApplicationResponse{
		Total:        total,
		Applications: applicationInfos,
	}, nil
}

func (s *ApplicationService) Get(ctx context.Context, req *pb.GetApplicationRequest) (*pb.ApplicationInfo, error) {
	application, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ApplicationInfo{
		Id:          application.Id,
		Name:        application.Name,
		Description: application.Description,
		AppKey:      application.AppKey,
		AppSecret:   application.AppSecret,
		CallbackUrl: application.CallbackUrl,
		CreatedAt:   application.CreatedAt.Unix(),
		UpdatedAt:   application.UpdatedAt.Unix(),
	}, nil
}

func (s *ApplicationService) Create(ctx context.Context, req *pb.CreateApplicationRequest) (*pb.ApplicationInfo, error) {
	application := &biz.Application{
		Name:        req.Name,
		Description: req.Description,
		CallbackUrl: req.CallbackUrl,
	}

	createdApplication, err := s.uc.Create(ctx, application)
	if err != nil {
		return nil, err
	}

	return &pb.ApplicationInfo{
		Id:          createdApplication.Id,
		Name:        createdApplication.Name,
		Description: createdApplication.Description,
		AppKey:      createdApplication.AppKey,
		AppSecret:   createdApplication.AppSecret,
		CallbackUrl: createdApplication.CallbackUrl,
		CreatedAt:   createdApplication.CreatedAt.Unix(),
		UpdatedAt:   createdApplication.UpdatedAt.Unix(),
	}, nil
}

func (s *ApplicationService) Update(ctx context.Context, req *pb.UpdateApplicationRequest) (*emptypb.Empty, error) {
	application := &biz.Application{
		ID:          biz.ID{Id: req.Id},
		Name:        req.Name,
		Description: req.Description,
		CallbackUrl: req.CallbackUrl,
	}

	err := s.uc.Update(ctx, application)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApplicationService) Delete(ctx context.Context, req *pb.DeleteApplicationRequest) (*emptypb.Empty, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
