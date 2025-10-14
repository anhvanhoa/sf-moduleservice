package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) DeleteRolePermission(ctx context.Context, req *proto_role_permission.DeleteRolePermissionRequest) (*proto_role_permission.DeleteRolePermissionResponse, error) {
	err := s.rolePermissionUsecase.Delete(ctx, req.RoleId, req.PermissionId)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.DeleteRolePermissionResponse{}, nil
}
