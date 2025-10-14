package role_permission_service

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) ListRolePermissions(ctx context.Context, req *proto_role_permission.ListRolePermissionsRequest) (*proto_role_permission.ListRolePermissionsResponse, error) {
	pagination := common.Pagination{Page: 1, PageSize: 10}
	if req.Pagination != nil {
		pagination.Page = int(req.Pagination.Page)
		pagination.PageSize = int(req.Pagination.PageSize)
	}
	var filter entity.RolePermissionFilter
	if req.Filter != nil {
		filter.RoleID = req.Filter.RoleId
		filter.PermissionID = req.Filter.PermissionId
	}
	result, err := s.rolePermissionUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoRolePermissions := make([]*proto_role_permission.RolePermission, len(result.Data))
	for i, rp := range result.Data {
		protoRolePermissions[i] = s.convertEntityToProtoRolePermission(rp)
	}

	return &proto_role_permission.ListRolePermissionsResponse{
		RolePermissions: protoRolePermissions,
		Pagination: &proto_common.PaginationResponse{
			Page:       int32(result.Page),
			PageSize:   int32(result.PageSize),
			TotalPages: int32(result.TotalPages),
			Total:      int32(result.Total),
		},
	}, nil
}
