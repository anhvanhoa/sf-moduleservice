package permission_service

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) ListPermissions(ctx context.Context, req *proto_permission.ListPermissionsRequest) (*proto_permission.ListPermissionsResponse, error) {
	pagination := common.Pagination{Page: 1, PageSize: 10}
	if req.Pagination != nil {
		pagination.Page = int(req.Pagination.Page)
		pagination.PageSize = int(req.Pagination.PageSize)
	}
	var filter entity.PermissionFilter
	if req.Filter != nil {
		filter.Resource = req.Filter.Resource
		filter.Action = req.Filter.Action
	}
	result, err := s.permissionUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoPermissions := make([]*proto_permission.Permission, len(result.Data))
	for i, permission := range result.Data {
		protoPermissions[i] = s.convertEntityToProtoPermission(permission)
	}

	return &proto_permission.ListPermissionsResponse{
		Permissions: protoPermissions,
		Pagination: &proto_common.PaginationResponse{
			Page:       int32(result.Page),
			PageSize:   int32(result.PageSize),
			TotalPages: int32(result.TotalPages),
			Total:      int32(result.Total),
		},
	}, nil
}
