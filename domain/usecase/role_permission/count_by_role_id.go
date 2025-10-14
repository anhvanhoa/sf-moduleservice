package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountByRoleIDUsecase interface {
	Execute(ctx context.Context, roleID string) (int64, error)
}

type CountByRoleIDUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewCountByRoleIDUsecase(rolePermissionRepository repository.RolePermissionRepository) CountByRoleIDUsecase {
	return &CountByRoleIDUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *CountByRoleIDUsecaseImpl) Execute(ctx context.Context, roleID string) (int64, error) {
	return u.rolePermissionRepository.CountByRoleID(ctx, roleID)
}
