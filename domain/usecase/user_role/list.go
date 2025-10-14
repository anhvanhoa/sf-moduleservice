package user_role

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type ListUserRolesUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error)
}

type ListUserRolesUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
}

func NewListUserRolesUsecase(userRoleRepository repository.UserRoleRepository) ListUserRolesUsecase {
	return &ListUserRolesUsecaseImpl{
		userRoleRepository: userRoleRepository,
	}
}

func (u *ListUserRolesUsecaseImpl) Execute(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error) {
	return u.userRoleRepository.List(ctx, pagination, filter)
}
