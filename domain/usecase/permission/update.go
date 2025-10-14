package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type UpdatePermissionUsecase interface {
	Execute(ctx context.Context, permission *entity.Permission) error
}

type UpdatePermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewUpdatePermissionUsecase(permissionRepository repository.PermissionRepository) UpdatePermissionUsecase {
	return &UpdatePermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *UpdatePermissionUsecaseImpl) Execute(ctx context.Context, permission *entity.Permission) error {
	return u.permissionRepository.Update(ctx, permission)
}
