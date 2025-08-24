package modulechildservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	"time"

	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) UpdateModuleChild(ctx context.Context, req *proto_module_child.UpdateModuleChildRequest) (*proto_module_child.UpdateModuleChildResponse, error) {
	moduleChild := &entity.ModuleChild{
		ID:        req.Id,
		ModuleID:  req.ModuleId,
		Name:      req.Name,
		Path:      req.Path,
		Method:    req.Method,
		IsPrivate: req.IsPrivate,
		Status:    common.Status(req.Status),
	}

	err := s.updateChildUc.Update(ctx, moduleChild)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Get the updated module child to return complete data
	updatedModuleChild, err := s.getChildUc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleChildProto := &proto_module_child.ModuleChild{
		Id:        updatedModuleChild.ID,
		ModuleId:  updatedModuleChild.ModuleID,
		Name:      updatedModuleChild.Name,
		Path:      updatedModuleChild.Path,
		Method:    updatedModuleChild.Method,
		IsPrivate: updatedModuleChild.IsPrivate,
		Status:    string(updatedModuleChild.Status),
		CreatedAt: updatedModuleChild.CreatedAt.Format(time.RFC3339),
	}

	if updatedModuleChild.UpdatedAt != nil {
		moduleChildProto.UpdatedAt = updatedModuleChild.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_module_child.UpdateModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
