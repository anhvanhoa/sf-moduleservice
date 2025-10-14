package resource_permission_service

import (
	"context"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) CountByUserID(ctx context.Context, req *proto_resource_permission.CountByUserIDRequest) (*proto_resource_permission.CountByUserIDResponse, error) {
	count, err := s.resourcePermissionUsecase.CountByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.CountByUserIDResponse{
		Count: count,
	}, nil
}
