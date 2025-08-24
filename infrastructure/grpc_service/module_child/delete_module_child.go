package modulechildservice

import (
	"context"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) DeleteModuleChild(ctx context.Context, req *proto_module.DeleteModuleChildRequest) (*proto_module.DeleteModuleChildResponse, error) {
	err := s.deleteChildUc.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto_module.DeleteModuleChildResponse{
		Success: true,
	}, nil
}
