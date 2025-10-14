package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteRolePermissionUsecase interface {
	Execute(ctx context.Context, roleID, permissionID string) error
}

type DeleteRolePermissionUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewDeleteRolePermissionUsecase(rolePermissionRepository repository.RolePermissionRepository) DeleteRolePermissionUsecase {
	return &DeleteRolePermissionUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *DeleteRolePermissionUsecaseImpl) Execute(ctx context.Context, roleID, permissionID string) error {
	err := u.rolePermissionRepository.Delete(ctx, roleID, permissionID)
	if err != nil {
		return ErrDeleteRolePermission
	}
	return nil
}
