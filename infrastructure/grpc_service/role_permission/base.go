package role_permission_service

import (
	"module-service/domain/entity"
	"module-service/domain/usecase/role_permission"
	"module-service/infrastructure/repo"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

type rolePermissionService struct {
	proto_role_permission.UnimplementedRolePermissionServiceServer
	rolePermissionUsecase role_permission.RolePermissionUsecaseI
}

func NewRolePermissionServer(repos repo.Repositories) proto_role_permission.RolePermissionServiceServer {
	rolePermissionRepo := repos.RolePermissionRepository()
	rolePermissionUC := role_permission.NewRolePermissionUsecase(rolePermissionRepo)
	return &rolePermissionService{
		rolePermissionUsecase: rolePermissionUC,
	}
}

func (s *rolePermissionService) convertEntityToProtoRolePermission(rp *entity.RolePermission) *proto_role_permission.RolePermission {
	return &proto_role_permission.RolePermission{
		RoleId:       rp.RoleID,
		PermissionId: rp.PermissionID,
	}
}
