package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListPermissionsUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) ([]*entity.Permission, int64, error)
}

type ListPermissionsUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewListPermissionsUsecase(permissionRepository repository.PermissionRepository) ListPermissionsUsecase {
	return &ListPermissionsUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *ListPermissionsUsecaseImpl) Execute(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) ([]*entity.Permission, int64, error) {
	return u.permissionRepository.List(ctx, pagination, filter)
}
