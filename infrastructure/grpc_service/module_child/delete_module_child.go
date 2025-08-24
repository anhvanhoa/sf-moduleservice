package modulechildservice

import (
	"context"

	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) DeleteModuleChild(ctx context.Context, req *proto_module_child.DeleteModuleChildRequest) (*proto_module_child.DeleteModuleChildResponse, error) {
	err := s.deleteChildUc.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto_module_child.DeleteModuleChildResponse{
		Success: true,
	}, nil
}
