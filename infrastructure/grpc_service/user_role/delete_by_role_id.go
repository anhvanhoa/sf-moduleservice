package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) DeleteByRoleID(ctx context.Context, req *proto_user_role.DeleteByRoleIDRequest) (*proto_user_role.DeleteByRoleIDResponse, error) {
	err := s.userRoleUsecase.DeleteByRoleID(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.DeleteByRoleIDResponse{}, nil
}
