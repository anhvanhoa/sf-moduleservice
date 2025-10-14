package role_permission_service

import (
	"context"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) DeleteByPermissionID(ctx context.Context, req *proto_role_permission.DeleteByPermissionIDRequest) (*proto_role_permission.DeleteByPermissionIDResponse, error) {
	err := s.rolePermissionUsecase.DeleteByPermissionID(ctx, req.PermissionId)
	if err != nil {
		return nil, err
	}
	return &proto_role_permission.DeleteByPermissionIDResponse{}, nil
}
