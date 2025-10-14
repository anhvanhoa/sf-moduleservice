package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetPermissionUsecase interface {
	Execute(ctx context.Context, id string) (*entity.Permission, error)
}

type GetPermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewGetPermissionUsecase(permissionRepository repository.PermissionRepository) GetPermissionUsecase {
	return &GetPermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *GetPermissionUsecaseImpl) Execute(ctx context.Context, id string) (*entity.Permission, error) {
	return u.permissionRepository.GetByID(ctx, id)
}
