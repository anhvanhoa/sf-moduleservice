package modulechildservice

import (
	"context"
	proto "module-service/proto/gen/module/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) GetModuleChild(ctx context.Context, req *proto.GetModuleChildRequest) (*proto.GetModuleChildResponse, error) {
	moduleChild, err := s.getChildUc.GetByID(ctx, req.Id)
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

	return &proto.GetModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
