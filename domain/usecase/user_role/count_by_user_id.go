package user_role

import (
	"context"
	"module-service/domain/repository"
)

type CountByUserIDUsecase interface {
	Execute(ctx context.Context, userID string) (int64, error)
}

type CountByUserIDUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewCountByUserIDUsecase(userRoleRepository repository.UserRoleRepository) CountByUserIDUsecase {
	return &CountByUserIDUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *CountByUserIDUsecaseImpl) Execute(ctx context.Context, userID string) (int64, error) {
	return u.userRoleRepository.CountByUserID(ctx, userID)
}
