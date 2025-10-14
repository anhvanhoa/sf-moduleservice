package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) DeleteByUserAndResource(ctx context.Context, req *proto_resource_permission.DeleteByUserAndResourceRequest) (*proto_resource_permission.DeleteByUserAndResourceResponse, error) {
	err := s.resourcePermissionUsecase.DeleteByUserAndResource(ctx, req.UserId, req.ResourceType, req.ResourceId)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.DeleteByUserAndResourceResponse{}, nil
}
