package user_role_service

import (
	"module-service/domain/entity"
	"module-service/domain/usecase/user_role"
	"module-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/utils"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

type userRoleService struct {
	proto_user_role.UnimplementedUserRoleServiceServer
	userRoleUsecase user_role.UserRoleUsecaseI
}

func NewUserRoleServer(repos repo.Repositories, helper utils.Helper) proto_user_role.UserRoleServiceServer {
	userRoleRepo := repos.UserRoleRepository()
	userRoleUC := user_role.NewUserRoleUsecase(userRoleRepo, helper)
	return &userRoleService{
		userRoleUsecase: userRoleUC,
	}
}

func (s *userRoleService) convertEntityToProtoUserRole(ur *entity.UserRole) *proto_user_role.UserRole {
	return &proto_user_role.UserRole{
		UserId: ur.UserID,
		RoleId: ur.RoleID,
	}
}
