package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) ExistsResourcePermission(ctx context.Context, req *proto_resource_permission.ExistsResourcePermissionRequest) (*proto_resource_permission.ExistsResourcePermissionResponse, error) {
	exists, err := s.resourcePermissionUsecase.Exists(ctx, req.UserId, req.ResourceType, req.ResourceId, req.Action)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.ExistsResourcePermissionResponse{
		Exists: exists,
	}, nil
}
