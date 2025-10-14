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
	count, err := u.rolePermissionRepository.Count(ctx)
	if err != nil {
		return 0, ErrCountRolePermissions
	}
	return count, nil
}
