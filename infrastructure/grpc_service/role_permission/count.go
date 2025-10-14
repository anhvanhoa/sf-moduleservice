package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) CountRolePermissions(ctx context.Context, req *proto_role_permission.CountRolePermissionsRequest) (*proto_role_permission.CountRolePermissionsResponse, error) {
	count, err := s.rolePermissionUsecase.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.CountRolePermissionsResponse{
		Count: count,
	}, nil
}
