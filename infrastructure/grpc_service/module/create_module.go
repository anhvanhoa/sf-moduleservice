package moduleservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	proto "module-service/proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) CreateModule(ctx context.Context, req *proto.CreateModuleRequest) (*proto.CreateModuleResponse, error) {
	module := entity.Module{
		Name:        req.Name,
		Description: req.Description,
		Status:      common.Status(req.Status),
	}
	err := s.createUc.CreateModule(ctx, &module)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CreateModuleResponse{
		Module: &proto.Module{
			Id:          module.ID,
			Name:        module.Name,
			Description: module.Description,
			Status:      string(module.Status),
		},
	}, nil
}
