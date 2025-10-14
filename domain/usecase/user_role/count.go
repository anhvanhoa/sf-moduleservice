package user_role

import (
	"context"
	"module-service/domain/repository"
)

type CountUserRolesUsecase interface {
	Execute(ctx context.Context) (int64, error)
}

type CountUserRolesUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewCountUserRolesUsecase(userRoleRepository repository.UserRoleRepository) CountUserRolesUsecase {
	return &CountUserRolesUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *CountUserRolesUsecaseImpl) Execute(ctx context.Context) (int64, error) {
	count, err := u.userRoleRepository.Count(ctx)
	if err != nil {
		return 0, ErrCountUserRoles
	}
	return count, nil
}
