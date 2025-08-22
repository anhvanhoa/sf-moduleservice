package moduleservice

import (
	"context"
	"module-service/domain/common"
	"time"
	proto "module-service/proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) ListModules(ctx context.Context, req *proto.ListModulesRequest) (*proto.ListModulesResponse, error) {
	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.Limit),
	}
	
	modules, total, err := s.listUc.List(ctx, pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	moduleProtos := make([]*proto.Module, len(modules))
	for i, module := range modules {
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
		
		moduleProtos[i] = moduleProto
	}
	
	totalPages := (int(total) + pagination.PageSize - 1) / pagination.PageSize
	
	return &proto.ListModulesResponse{
		Modules: moduleProtos,
		Pagination: &proto.PaginationResponse{
			Page:       int32(pagination.Page),
			Limit:      int32(pagination.PageSize),
			Total:      int32(total),
			TotalPages: int32(totalPages),
		},
	}, nil
}
