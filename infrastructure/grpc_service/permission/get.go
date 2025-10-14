package permission_service

import (
	"context"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) GetPermission(ctx context.Context, req *proto_permission.GetPermissionRequest) (*proto_permission.GetPermissionResponse, error) {
	permission, err := s.permissionUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_permission.GetPermissionResponse{
		Permission: s.convertEntityToProtoPermission(permission),
	}, nil
}
