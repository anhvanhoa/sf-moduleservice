package modulechildservice

import (
	"context"
	proto "module-service/proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleChildService) DeleteModuleChild(ctx context.Context, req *proto.DeleteModuleChildRequest) (*proto.DeleteModuleChildResponse, error) {
	err := s.deleteChildUc.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteModuleChildResponse{
		Success: true,
	}, nil
}
