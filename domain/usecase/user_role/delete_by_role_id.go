package user_role

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByRoleIDUsecase interface {
	Execute(ctx context.Context, roleID string) error
}

type DeleteByRoleIDUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewDeleteByRoleIDUsecase(userRoleRepository repository.UserRoleRepository) DeleteByRoleIDUsecase {
	return &DeleteByRoleIDUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *DeleteByRoleIDUsecaseImpl) Execute(ctx context.Context, roleID string) error {
	return u.userRoleRepository.DeleteByRoleID(ctx, roleID)
}
