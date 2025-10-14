package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListResourcePermissionsUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) ([]*entity.ResourcePermission, int64, error)
}

type ListResourcePermissionsUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewListResourcePermissionsUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) ListResourcePermissionsUsecase {
	return &ListResourcePermissionsUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *ListResourcePermissionsUsecaseImpl) Execute(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) ([]*entity.ResourcePermission, int64, error) {
	return u.resourcePermissionRepository.List(ctx, pagination, filter)
}
