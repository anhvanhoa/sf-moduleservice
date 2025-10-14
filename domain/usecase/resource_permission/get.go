package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetResourcePermissionUsecase interface {
	Execute(ctx context.Context, id string) (*entity.ResourcePermission, error)
}

type GetResourcePermissionUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewGetResourcePermissionUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) GetResourcePermissionUsecase {
	return &GetResourcePermissionUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *GetResourcePermissionUsecaseImpl) Execute(ctx context.Context, id string) (*entity.ResourcePermission, error) {
	return u.resourcePermissionRepository.GetByID(ctx, id)
}
