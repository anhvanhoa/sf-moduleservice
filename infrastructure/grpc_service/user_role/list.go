package user_role_service

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) ListUserRoles(ctx context.Context, req *proto_user_role.ListUserRolesRequest) (*proto_user_role.ListUserRolesResponse, error) {
	pagination := common.Pagination{Page: 1, PageSize: 10}
	if req.Pagination != nil {
		pagination.Page = int(req.Pagination.Page)
		pagination.PageSize = int(req.Pagination.PageSize)
	}
	var filter entity.UserRoleFilter
	if req.Filter != nil {
		filter.UserID = req.Filter.UserId
		filter.RoleID = req.Filter.RoleId
	}
	result, err := s.userRoleUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoUserRoles := make([]*proto_user_role.UserRole, len(result.Data))
	for i, ur := range result.Data {
		protoUserRoles[i] = s.convertEntityToProtoUserRole(ur)
	}

	return &proto_user_role.ListUserRolesResponse{
		UserRoles: protoUserRoles,
		Pagination: &proto_common.PaginationResponse{
			Page:       int32(result.Page),
			PageSize:   int32(result.PageSize),
			TotalPages: int32(result.TotalPages),
			Total:      int32(result.Total),
		},
	}, nil
}
