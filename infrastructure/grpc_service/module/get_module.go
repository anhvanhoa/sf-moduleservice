package moduleservice

import (
	"context"
	"time"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) GetModule(ctx context.Context, req *proto_module.GetModuleRequest) (*proto_module.GetModuleResponse, error) {
	module, err := s.getUc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleProto := &proto_module.Module{
		Id:          module.ID,
		Name:        module.Name,
		Description: module.Description,
		Status:      string(module.Status),
		CreatedAt:   module.CreatedAt.Format(time.RFC3339),
	}

	if module.UpdatedAt != nil {
		moduleProto.UpdatedAt = module.UpdatedAt.Format(time.RFC3339)
	}

	return &proto_module.GetModuleResponse{
		Module: moduleProto,
	}, nil
}
