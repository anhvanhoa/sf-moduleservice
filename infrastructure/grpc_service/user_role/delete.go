package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) DeleteUserRole(ctx context.Context, req *proto_user_role.DeleteUserRoleRequest) (*proto_user_role.DeleteUserRoleResponse, error) {
	err := s.userRoleUsecase.Delete(ctx, req.UserId, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.DeleteUserRoleResponse{}, nil
}
