package user_role

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateUserRoleUsecase interface {
	Execute(ctx context.Context, userRole *entity.UserRole) error
}

type CreateUserRoleUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewCreateUserRoleUsecase(userRoleRepository repository.UserRoleRepository) CreateUserRoleUsecase {
	return &CreateUserRoleUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *CreateUserRoleUsecaseImpl) Execute(ctx context.Context, userRole *entity.UserRole) error {
	return u.userRoleRepository.Create(ctx, userRole)
}
