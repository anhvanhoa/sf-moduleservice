package modulechildservice

import (
	"context"
	"time"

	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) GetModuleChild(ctx context.Context, req *proto_module_child.GetModuleChildRequest) (*proto_module_child.GetModuleChildResponse, error) {
	moduleChild, err := s.getChildUc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	moduleChildProto := &proto_module_child.ModuleChild{
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

	return &proto_module_child.GetModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
