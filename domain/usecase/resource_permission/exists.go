package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type ExistsResourcePermissionUsecase interface {
	Execute(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error)
}

type ExistsResourcePermissionUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewExistsResourcePermissionUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) ExistsResourcePermissionUsecase {
	return &ExistsResourcePermissionUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *ExistsResourcePermissionUsecaseImpl) Execute(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error) {
	isExists, err := u.resourcePermissionRepository.Exists(ctx, userID, resourceType, resourceID, action)
	if err != nil {
		return false, ErrExistsResourcePermission
	}
	return isExists, nil
}
