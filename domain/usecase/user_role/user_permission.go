package user_role

import (
	"context"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/user_context"
)

type GetUserPermissionsUsecase interface {
	Execute(ctx context.Context, userID string) (user_context.UserContext, error)
}

type GetUserPermissionsUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewGetUserPermissionsUsecase(userRoleRepository repository.UserRoleRepository) GetUserPermissionsUsecase {
	return &GetUserPermissionsUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *GetUserPermissionsUsecaseImpl) Execute(ctx context.Context, userID string) (user_context.UserContext, error) {
	userPermission, err := u.userRoleRepository.GetUserPermissions(ctx, userID)
	if err != nil {
		return user_context.UserContext{}, err
	}
	return userPermission, nil
}
