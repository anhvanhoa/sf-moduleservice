package moduleservice

import (
	"context"
	proto "module-service/proto/gen/module/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *moduleService) DeleteModule(ctx context.Context, req *proto.DeleteModuleRequest) (*proto.DeleteModuleResponse, error) {
	err := s.deleteUc.DeleteByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &proto.DeleteModuleResponse{
		Success: true,
	}, nil
}
