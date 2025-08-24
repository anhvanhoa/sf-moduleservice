package modulechildservice

import (
	"context"
	"module-service/domain/common"
	"time"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) ListModuleChildren(ctx context.Context, req *proto_module.ListModuleChildrenRequest) (*proto_module.ListModuleChildrenResponse, error) {
	pagination := &common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.Limit),
	}

	moduleChildren, total, err := s.listChildrenUc.List(ctx, pagination, req.ModuleId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleChildProtos := make([]*proto_module.ModuleChild, len(moduleChildren))
	for i, moduleChild := range moduleChildren {
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

		moduleChildProtos[i] = moduleChildProto
	}

	totalPages := (int(total) + pagination.PageSize - 1) / pagination.PageSize

	return &proto_module.ListModuleChildrenResponse{
		ModuleChildren: moduleChildProtos,
		Pagination: &proto_module.PaginationResponse{
			Page:       int32(pagination.Page),
			Limit:      int32(pagination.PageSize),
			Total:      int32(total),
			TotalPages: int32(totalPages),
		},
	}, nil
}
