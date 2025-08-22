package moduleservice

import (
	"context"
	"module-service/domain/common"
	"module-service/domain/entity"
	"time"
	proto "module-service/proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) UpdateModule(ctx context.Context, req *proto.UpdateModuleRequest) (*proto.UpdateModuleResponse, error) {
	module := &entity.Module{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Status:      common.Status(req.Status),
	}
	
	err := s.updateUc.Update(ctx, module)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	// Get the updated module to return complete data
	updatedModule, err := s.getUc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	moduleProto := &proto.Module{
		Id:          updatedModule.ID,
		Name:        updatedModule.Name,
		Description: updatedModule.Description,
		Status:      string(updatedModule.Status),
		CreatedAt:   updatedModule.CreatedAt.Format(time.RFC3339),
	}
	
	if updatedModule.UpdatedAt != nil {
		moduleProto.UpdatedAt = updatedModule.UpdatedAt.Format(time.RFC3339)
	}
	
	return &proto.UpdateModuleResponse{
		Module: moduleProto,
	}, nil
}
