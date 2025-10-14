package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByResourceUsecase interface {
	Execute(ctx context.Context, resourceType, resourceID string) error
}

type DeleteByResourceUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewDeleteByResourceUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) DeleteByResourceUsecase {
	return &DeleteByResourceUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *DeleteByResourceUsecaseImpl) Execute(ctx context.Context, resourceType, resourceID string) error {
	return u.resourcePermissionRepository.DeleteByResource(ctx, resourceType, resourceID)
}
