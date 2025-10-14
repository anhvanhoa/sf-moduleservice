package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type ExistsRolePermissionUsecase interface {
	Execute(ctx context.Context, roleID, permissionID string) (bool, error)
}

type ExistsRolePermissionUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewExistsRolePermissionUsecase(rolePermissionRepository repository.RolePermissionRepository) ExistsRolePermissionUsecase {
	return &ExistsRolePermissionUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *ExistsRolePermissionUsecaseImpl) Execute(ctx context.Context, roleID, permissionID string) (bool, error) {
	exists, err := u.rolePermissionRepository.Exists(ctx, roleID, permissionID)
	if err != nil {
		return false, ErrExistsRolePermission
	}
	return exists, nil
}
