package permission

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByResourceAndActionUsecase interface {
	Execute(ctx context.Context, resource, action string) error
}

type DeleteByResourceAndActionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewDeleteByResourceAndActionUsecase(permissionRepository repository.PermissionRepository) DeleteByResourceAndActionUsecase {
	return &DeleteByResourceAndActionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *DeleteByResourceAndActionUsecaseImpl) Execute(ctx context.Context, resource, action string) error {
	err := u.permissionRepository.DeleteByResourceAndAction(ctx, resource, action)
	if err != nil {
		return ErrDeleteByResourceAndAction
	}
	return nil
}
