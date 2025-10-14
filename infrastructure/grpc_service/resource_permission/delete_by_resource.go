package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) DeleteByResource(ctx context.Context, req *proto_resource_permission.DeleteByResourceRequest) (*proto_resource_permission.DeleteByResourceResponse, error) {
	err := s.resourcePermissionUsecase.DeleteByResource(ctx, req.ResourceType, req.ResourceId)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.DeleteByResourceResponse{}, nil
}
