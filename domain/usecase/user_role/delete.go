package user_role

import (
	"context"
	"module-service/domain/repository"
)

type DeleteUserRoleUsecase interface {
	Execute(ctx context.Context, userID, roleID string) error
}

type DeleteUserRoleUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewDeleteUserRoleUsecase(userRoleRepository repository.UserRoleRepository) DeleteUserRoleUsecase {
	return &DeleteUserRoleUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *DeleteUserRoleUsecaseImpl) Execute(ctx context.Context, userID, roleID string) error {
	return u.userRoleRepository.Delete(ctx, userID, roleID)
}
