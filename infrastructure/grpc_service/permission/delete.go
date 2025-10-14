package permission_service

import (
	"context"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) DeletePermission(ctx context.Context, req *proto_permission.DeletePermissionRequest) (*proto_permission.DeletePermissionResponse, error) {
	err := s.permissionUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_permission.DeletePermissionResponse{}, nil
}
