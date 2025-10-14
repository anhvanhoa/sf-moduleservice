package user_role

import (
	"context"
	"module-service/domain/repository"
)

type ExistsUserRoleUsecase interface {
	Execute(ctx context.Context, userID, roleID string) (bool, error)
}

type ExistsUserRoleUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewExistsUserRoleUsecase(userRoleRepository repository.UserRoleRepository) ExistsUserRoleUsecase {
	return &ExistsUserRoleUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *ExistsUserRoleUsecaseImpl) Execute(ctx context.Context, userID, roleID string) (bool, error) {
	return u.userRoleRepository.Exists(ctx, userID, roleID)
}
