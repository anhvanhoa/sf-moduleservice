package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByUserIDUsecase interface {
	Execute(ctx context.Context, userID string) error
}

type DeleteByUserIDUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewDeleteByUserIDUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) DeleteByUserIDUsecase {
	return &DeleteByUserIDUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *DeleteByUserIDUsecaseImpl) Execute(ctx context.Context, userID string) error {
	return u.resourcePermissionRepository.DeleteByUserID(ctx, userID)
}
