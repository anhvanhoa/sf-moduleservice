package modulechildservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	"time"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) CreateModuleChild(ctx context.Context, req *proto_module.CreateModuleChildRequest) (*proto_module.CreateModuleChildResponse, error) {
	moduleChild := &entity.ModuleChild{
		ModuleID:  req.ModuleId,
		Name:      req.Name,
		Path:      req.Path,
		Method:    req.Method,
		IsPrivate: req.IsPrivate,
		Status:    common.Status(req.Status),
	}

	err := s.createChildUc.CreateModuleChild(ctx, moduleChild)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleChildProto := &proto_module.ModuleChild{
		Id:        moduleChild.ID,
		ModuleId:  moduleChild.ModuleID,
		Name:      moduleChild.Name,
		Path:      moduleChild.Path,
		Method:    moduleChild.Method,
		IsPrivate: moduleChild.IsPrivate,
		Status:    string(moduleChild.Status),
		CreatedAt: moduleChild.CreatedAt.Format(time.RFC3339),
	}

	if moduleChild.UpdatedAt != nil {
		moduleChildProto.UpdatedAt = moduleChild.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_module.CreateModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
