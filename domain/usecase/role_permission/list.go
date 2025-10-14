package role_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListRolePermissionsUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) ([]*entity.RolePermission, int64, error)
}

type ListRolePermissionsUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewListRolePermissionsUsecase(rolePermissionRepository repository.RolePermissionRepository) ListRolePermissionsUsecase {
	return &ListRolePermissionsUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *ListRolePermissionsUsecaseImpl) Execute(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) ([]*entity.RolePermission, int64, error) {
	return u.rolePermissionRepository.List(ctx, pagination, filter)
}
