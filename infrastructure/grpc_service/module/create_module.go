package moduleservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) CreateModule(ctx context.Context, req *proto_module.CreateModuleRequest) (*proto_module.CreateModuleResponse, error) {
	module := entity.Module{
		Name:        req.Name,
		Description: req.Description,
		Status:      common.Status(req.Status),
	}
	err := s.createUc.CreateModule(ctx, &module)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_module.CreateModuleResponse{
		Module: &proto_module.Module{
			Id:          module.ID,
			Name:        module.Name,
			Description: module.Description,
			Status:      string(module.Status),
		},
	}, nil
}
