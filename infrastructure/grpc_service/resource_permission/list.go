package resource_permission_service

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) ListResourcePermissions(ctx context.Context, req *proto_resource_permission.ListResourcePermissionsRequest) (*proto_resource_permission.ListResourcePermissionsResponse, error) {
	pagination := common.Pagination{Page: 1, PageSize: 10}
	if req.Pagination != nil {
		pagination.Page = int(req.Pagination.Page)
		pagination.PageSize = int(req.Pagination.PageSize)
	}
	var filter entity.ResourcePermissionFilter
	if req.Filter != nil {
		filter.UserID = req.Filter.UserId
		filter.ResourceType = req.Filter.ResourceType
		filter.ResourceID = req.Filter.ResourceId
		filter.Action = req.Filter.Action
	}
	resourcePermissions, _, err := s.resourcePermissionUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoResourcePermissions := make([]*proto_resource_permission.ResourcePermission, len(resourcePermissions))
	for i, rp := range resourcePermissions {
		protoResourcePermissions[i] = s.convertEntityToProtoResourcePermission(rp)
	}

	return &proto_resource_permission.ListResourcePermissionsResponse{
		ResourcePermissions: protoResourcePermissions,
	}, nil
}
