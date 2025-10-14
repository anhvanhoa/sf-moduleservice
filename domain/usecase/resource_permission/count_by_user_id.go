package resource_permission

import (
	"context"
	"module-service/domain/repository"
)

type CountByUserIDUsecase interface {
	Execute(ctx context.Context, userID string) (int64, error)
}

type CountByUserIDUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewCountByUserIDUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) CountByUserIDUsecase {
	return &CountByUserIDUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *CountByUserIDUsecaseImpl) Execute(ctx context.Context, userID string) (int64, error) {
	return u.resourcePermissionRepository.CountByUserID(ctx, userID)
}
