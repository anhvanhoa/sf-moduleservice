package role_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateRolePermissionUsecase interface {
	Execute(ctx context.Context, rolePermission *entity.RolePermission) error
}

type CreateRolePermissionUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewCreateRolePermissionUsecase(rolePermissionRepository repository.RolePermissionRepository) CreateRolePermissionUsecase {
	return &CreateRolePermissionUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *CreateRolePermissionUsecaseImpl) Execute(ctx context.Context, rolePermission *entity.RolePermission) error {
	err := u.rolePermissionRepository.Create(ctx, rolePermission)
	if err != nil {
		return ErrCreateRolePermission
	}
	return nil
}
