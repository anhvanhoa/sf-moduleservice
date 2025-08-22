package moduleservice

import (
	"context"
	proto "module-service/proto/gen/module/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) GetModule(ctx context.Context, req *proto.GetModuleRequest) (*proto.GetModuleResponse, error) {
	module, err := s.getUc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleProto := &proto.Module{
		Id:          module.ID,
		Name:        module.Name,
		Description: module.Description,
		Status:      string(module.Status),
		CreatedAt:   module.CreatedAt.Format(time.RFC3339),
	}

	if module.UpdatedAt != nil {
		moduleProto.UpdatedAt = module.UpdatedAt.Format(time.RFC3339)
	}

	return &proto.GetModuleResponse{
		Module: moduleProto,
	}, nil
}
