package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) GetResourcePermission(ctx context.Context, req *proto_resource_permission.GetResourcePermissionRequest) (*proto_resource_permission.GetResourcePermissionResponse, error) {
	resourcePermission, err := s.resourcePermissionUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.GetResourcePermissionResponse{
		ResourcePermission: s.convertEntityToProtoResourcePermission(resourcePermission),
	}, nil
}
