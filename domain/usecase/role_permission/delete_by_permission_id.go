package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByPermissionIDUsecase interface {
	Execute(ctx context.Context, permissionID string) error
}

type DeleteByPermissionIDUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewDeleteByPermissionIDUsecase(rolePermissionRepository repository.RolePermissionRepository) DeleteByPermissionIDUsecase {
	return &DeleteByPermissionIDUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *DeleteByPermissionIDUsecaseImpl) Execute(ctx context.Context, permissionID string) error {
	return u.rolePermissionRepository.DeleteByPermissionID(ctx, permissionID)
}
