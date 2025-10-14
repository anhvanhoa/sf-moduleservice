package user_role

import (
	"context"
	"module-service/domain/repository"
)

type CountByRoleIDUsecase interface {
	Execute(ctx context.Context, roleID string) (int64, error)
}

type CountByRoleIDUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewCountByRoleIDUsecase(userRoleRepository repository.UserRoleRepository) CountByRoleIDUsecase {
	return &CountByRoleIDUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *CountByRoleIDUsecaseImpl) Execute(ctx context.Context, roleID string) (int64, error) {
	count, err := u.userRoleRepository.CountByRoleID(ctx, roleID)
	if err != nil {
		return 0, ErrCountByRoleID
	}
	return count, nil
}
