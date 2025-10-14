package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) CountByRoleID(ctx context.Context, req *proto_role_permission.CountByRoleIDRequest) (*proto_role_permission.CountByRoleIDResponse, error) {
	count, err := s.rolePermissionUsecase.CountByRoleID(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.CountByRoleIDResponse{
		Count: count,
	}, nil
}
