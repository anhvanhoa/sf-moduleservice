package moduleservice

import (
	"context"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) DeleteModule(ctx context.Context, req *proto_module.DeleteModuleRequest) (*proto_module.DeleteModuleResponse, error) {
	err := s.deleteUc.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto_module.DeleteModuleResponse{
		Success: true,
	}, nil
}
