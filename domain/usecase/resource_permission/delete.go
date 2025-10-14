package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteResourcePermissionUsecase interface {
	Execute(ctx context.Context, id string) error
}

type DeleteResourcePermissionUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewDeleteResourcePermissionUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) DeleteResourcePermissionUsecase {
	return &DeleteResourcePermissionUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *DeleteResourcePermissionUsecaseImpl) Execute(ctx context.Context, id string) error {
	return u.resourcePermissionRepository.Delete(ctx, id)
}
