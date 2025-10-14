package role_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountByPermissionIDUsecase interface {
	Execute(ctx context.Context, permissionID string) (int64, error)
}

type CountByPermissionIDUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
}

func NewCountByPermissionIDUsecase(rolePermissionRepository repository.RolePermissionRepository) CountByPermissionIDUsecase {
	return &CountByPermissionIDUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
	}
}

func (u *CountByPermissionIDUsecaseImpl) Execute(ctx context.Context, permissionID string) (int64, error) {
	count, err := u.rolePermissionRepository.CountByPermissionID(ctx, permissionID)
	if err != nil {
		return 0, ErrCountByPermissionID
	}
	return count, nil
}
