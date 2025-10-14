package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateResourcePermissionUsecase interface {
	Execute(ctx context.Context, resourcePermission *entity.ResourcePermission) error
}

type CreateResourcePermissionUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewCreateResourcePermissionUsecase(resourcePermissionRepository repository.ResourcePermissionRepository) CreateResourcePermissionUsecase {
	return &CreateResourcePermissionUsecaseImpl{
		resourcePermissionRepository: resourcePermissionRepository,
	}
}

func (u *CreateResourcePermissionUsecaseImpl) Execute(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	err := u.resourcePermissionRepository.Create(ctx, resourcePermission)
	if err != nil {
		return ErrCreateResourcePermission
	}
	return nil
}
