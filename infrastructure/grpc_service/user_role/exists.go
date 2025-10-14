package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) ExistsUserRole(ctx context.Context, req *proto_user_role.ExistsUserRoleRequest) (*proto_user_role.ExistsUserRoleResponse, error) {
	exists, err := s.userRoleUsecase.Exists(ctx, req.UserId, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.ExistsUserRoleResponse{
		Exists: exists,
	}, nil
}
