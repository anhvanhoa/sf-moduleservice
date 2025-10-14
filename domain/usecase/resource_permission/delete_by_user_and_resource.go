package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByUserAndResourceUsecase interface {
	Execute(ctx context.Context, userID, resourceType, resourceID string) error
}

type DeleteByUserAndResourceUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewDeleteByUserAndResourceUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) DeleteByUserAndResourceUsecase {
	return &DeleteByUserAndResourceUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *DeleteByUserAndResourceUsecaseImpl) Execute(ctx context.Context, userID, resourceType, resourceID string) error {
	return u.resourcePermissionRepository.DeleteByUserAndResource(ctx, userID, resourceType, resourceID)
}
