package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreatePermissionUsecase interface {
	Execute(ctx context.Context, permission *entity.Permission) error
}

type CreatePermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewCreatePermissionUsecase(permissionRepository repository.PermissionRepository) CreatePermissionUsecase {
	return &CreatePermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *CreatePermissionUsecaseImpl) Execute(ctx context.Context, permission *entity.Permission) error {
	return u.permissionRepository.Create(ctx, permission)
}
