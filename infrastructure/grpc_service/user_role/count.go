package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) CountUserRoles(ctx context.Context, req *proto_user_role.CountUserRolesRequest) (*proto_user_role.CountUserRolesResponse, error) {
	count, err := s.userRoleUsecase.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.CountUserRolesResponse{
		Count: count,
	}, nil
}
