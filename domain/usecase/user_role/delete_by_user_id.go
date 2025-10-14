package user_role

import (
	"context"
	"module-service/domain/repository"
)

type DeleteByUserIDUsecase interface {
	Execute(ctx context.Context, userID string) error
}

type DeleteByUserIDUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewDeleteByUserIDUsecase(userRoleRepository repository.UserRoleRepository) DeleteByUserIDUsecase {
	return &DeleteByUserIDUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *DeleteByUserIDUsecaseImpl) Execute(ctx context.Context, userID string) error {
	return u.userRoleRepository.DeleteByUserID(ctx, userID)
}
