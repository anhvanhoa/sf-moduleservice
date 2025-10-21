package permission_service

import (
	"module-service/domain/entity"
	"module-service/domain/usecase/permission"
	"module-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/utils"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

type permissionService struct {
	proto_permission.UnimplementedPermissionServiceServer
	permissionUsecase permission.PermissionUsecaseI
}

func NewPermissionServer(repos repo.Repositories, cacher cache.CacheI, helper utils.Helper) proto_permission.PermissionServiceServer {
	permissionRepo := repos.PermissionRepository()
	permissionUC := permission.NewPermissionUsecase(permissionRepo, cacher, helper)
	return &permissionService{
		permissionUsecase: permissionUC,
	}
}

func (s *permissionService) convertEntityToProtoPermission(permission *entity.Permission) *proto_permission.Permission {
	return &proto_permission.Permission{
		Id:          permission.ID,
		Resource:    permission.Resource,
		Action:      permission.Action,
		Description: permission.Description,
		IsPublic:    permission.IsPublic,
	}
}
