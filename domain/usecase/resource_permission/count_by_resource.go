package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountByResourceUsecase interface {
	Execute(ctx context.Context, resourceType, resourceID string) (int64, error)
}

type CountByResourceUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewCountByResourceUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) CountByResourceUsecase {
	return &CountByResourceUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *CountByResourceUsecaseImpl) Execute(ctx context.Context, resourceType, resourceID string) (int64, error) {
	count, err := u.resourcePermissionRepository.CountByResource(ctx, resourceType, resourceID)
	if err != nil {
		return 0, ErrCountByResource
	}
	return count, nil
}
