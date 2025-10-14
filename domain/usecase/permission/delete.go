package permission

import (
	"context"
	"module-service/domain/repository"
)

type DeletePermissionUsecase interface {
	Execute(ctx context.Context, id string) error
}

type DeletePermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewDeletePermissionUsecase(permissionRepository repository.PermissionRepository) DeletePermissionUsecase {
	return &DeletePermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *DeletePermissionUsecaseImpl) Execute(ctx context.Context, id string) error {
	return u.permissionRepository.Delete(ctx, id)
}
