package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateManyPermissionUsecase interface {
	Execute(ctx context.Context, permissions []*entity.Permission) error
}

type CreateManyPermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewCreateManyPermissionUsecase(permissionRepository repository.PermissionRepository) CreateManyPermissionUsecase {
	return &CreateManyPermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *CreateManyPermissionUsecaseImpl) Execute(ctx context.Context, permissions []*entity.Permission) error {
	err := u.permissionRepository.CreateMany(ctx, permissions)
	if err != nil {
		return ErrCreateManyPermission
	}
	return nil
}
