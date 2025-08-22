package modulechildservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	proto "module-service/proto/gen/module/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) CreateModuleChild(ctx context.Context, req *proto.CreateModuleChildRequest) (*proto.CreateModuleChildResponse, error) {
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

	moduleChildProto := &proto.ModuleChild{
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

	return &proto.CreateModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
