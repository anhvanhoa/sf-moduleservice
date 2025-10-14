package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) CountByPermissionID(ctx context.Context, req *proto_role_permission.CountByPermissionIDRequest) (*proto_role_permission.CountByPermissionIDResponse, error) {
	count, err := s.rolePermissionUsecase.CountByPermissionID(ctx, req.PermissionId)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.CountByPermissionIDResponse{
		Count: count,
	}, nil
}
