package permission_service

import (
	"context"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) DeleteByResourceAndAction(ctx context.Context, req *proto_permission.DeleteByResourceAndActionRequest) (*proto_permission.DeleteByResourceAndActionResponse, error) {
	err := s.permissionUsecase.DeleteByResourceAndAction(ctx, req.Resource, req.Action)
	if err != nil {
		return nil, err
	}
	return &proto_permission.DeleteByResourceAndActionResponse{}, nil
}
