package user_role_service

import (
	"context"
	"module-service/domain/entity"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) CreateUserRole(ctx context.Context, req *proto_user_role.CreateUserRoleRequest) (*proto_user_role.CreateUserRoleResponse, error) {
	userRole := s.convertRequestCreateToEntity(req)
	err := s.userRoleUsecase.Create(ctx, userRole)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.CreateUserRoleResponse{
		Data: s.convertEntityToProtoUserRole(userRole),
	}, nil
}

func (s *userRoleService) convertRequestCreateToEntity(req *proto_user_role.CreateUserRoleRequest) *entity.UserRole {
	return &entity.UserRole{
		UserID: req.UserId,
		RoleID: req.RoleId,
	}
}
