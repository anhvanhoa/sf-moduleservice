package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) CountByResource(ctx context.Context, req *proto_resource_permission.CountByResourceRequest) (*proto_resource_permission.CountByResourceResponse, error) {
	count, err := s.resourcePermissionUsecase.CountByResource(ctx, req.ResourceType, req.ResourceId)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.CountByResourceResponse{
		Count: count,
	}, nil
}
