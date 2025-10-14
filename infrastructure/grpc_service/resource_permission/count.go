package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) CountResourcePermissions(ctx context.Context, req *proto_resource_permission.CountResourcePermissionsRequest) (*proto_resource_permission.CountResourcePermissionsResponse, error) {
	count, err := s.resourcePermissionUsecase.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.CountResourcePermissionsResponse{
		Count: count,
	}, nil
}
