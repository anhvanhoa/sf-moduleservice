package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountRolePermissionsUsecase interface {
	Execute(ctx context.Context) (int64, error)
}

type CountRolePermissionsUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewCountRolePermissionsUsecase(rolePermissionRepository repository.RolePermissionRepository) CountRolePermissionsUsecase {
	return &CountRolePermissionsUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *CountRolePermissionsUsecaseImpl) Execute(ctx context.Context) (int64, error) {
	return u.rolePermissionRepository.Count(ctx)
}
