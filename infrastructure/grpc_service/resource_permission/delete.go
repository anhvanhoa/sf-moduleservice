package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) DeleteResourcePermission(ctx context.Context, req *proto_resource_permission.DeleteResourcePermissionRequest) (*proto_resource_permission.DeleteResourcePermissionResponse, error) {
	err := s.resourcePermissionUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.DeleteResourcePermissionResponse{}, nil
}
