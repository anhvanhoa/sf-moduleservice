package permission_service

import (
	"context"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) CountByResource(ctx context.Context, req *proto_permission.CountByResourceRequest) (*proto_permission.CountByResourceResponse, error) {
	count, err := s.permissionUsecase.CountByResource(ctx, req.Resource)
	if err != nil {
		return nil, err
	}
	return &proto_permission.CountByResourceResponse{
		Count: count,
	}, nil
}
