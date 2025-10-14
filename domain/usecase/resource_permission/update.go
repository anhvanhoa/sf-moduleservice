package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type UpdateResourcePermissionUsecase interface {
	Execute(ctx context.Context, resourcePermission *entity.ResourcePermission) error
}

type UpdateResourcePermissionUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewUpdateResourcePermissionUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) UpdateResourcePermissionUsecase {
	return &UpdateResourcePermissionUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *UpdateResourcePermissionUsecaseImpl) Execute(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	return u.resourcePermissionRepository.Update(ctx, resourcePermission)
}
