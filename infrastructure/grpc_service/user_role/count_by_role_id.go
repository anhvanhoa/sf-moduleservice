package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) CountByRoleID(ctx context.Context, req *proto_user_role.CountByRoleIDRequest) (*proto_user_role.CountByRoleIDResponse, error) {
	count, err := s.userRoleUsecase.CountByRoleID(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.CountByRoleIDResponse{
		Count: count,
	}, nil
}
