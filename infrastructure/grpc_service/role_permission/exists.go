package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) ExistsRolePermission(ctx context.Context, req *proto_role_permission.ExistsRolePermissionRequest) (*proto_role_permission.ExistsRolePermissionResponse, error) {
	exists, err := s.rolePermissionUsecase.Exists(ctx, req.RoleId, req.PermissionId)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.ExistsRolePermissionResponse{
		Exists: exists,
	}, nil
}
