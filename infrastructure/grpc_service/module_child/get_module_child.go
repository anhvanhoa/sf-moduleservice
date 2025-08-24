package modulechildservice

import (
	"context"
	"time"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) GetModuleChild(ctx context.Context, req *proto_module.GetModuleChildRequest) (*proto_module.GetModuleChildResponse, error) {
	moduleChild, err := s.getChildUc.GetByID(ctx, req.Id)
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

	return &proto_module.GetModuleChildResponse{
		ModuleChild: moduleChildProto,
	}, nil
}
