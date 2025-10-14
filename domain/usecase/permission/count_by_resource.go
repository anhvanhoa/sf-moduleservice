package permission

import (
	"context"
	"module-service/domain/repository"
)

type CountByResourceUsecase interface {
	Execute(ctx context.Context, resource string) (int64, error)
}

type CountByResourceUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewCountByResourceUsecase(permissionRepository repository.PermissionRepository) CountByResourceUsecase {
	return &CountByResourceUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *CountByResourceUsecaseImpl) Execute(ctx context.Context, resource string) (int64, error) {
	count, err := u.permissionRepository.CountByResource(ctx, resource)
	if err != nil {
		return 0, ErrCountByResource
	}
	return count, nil
}
