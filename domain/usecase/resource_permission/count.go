package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountResourcePermissionsUsecase interface {
	Execute(ctx context.Context) (int64, error)
}

type CountResourcePermissionsUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewCountResourcePermissionsUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) CountResourcePermissionsUsecase {
	return &CountResourcePermissionsUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *CountResourcePermissionsUsecaseImpl) Execute(ctx context.Context) (int64, error) {
	count, err := u.resourcePermissionRepository.Count(ctx)
	if err != nil {
		return 0, ErrCountResourcePermissions
	}
	return count, nil
}
