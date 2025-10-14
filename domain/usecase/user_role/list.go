package user_role

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListUserRolesUsecase interface {
	Execute(
		ctx context.Context,
		pagination common.Pagination,
		filter entity.UserRoleFilter,
	) (common.PaginationResult[*entity.UserRole], error)
}

type ListUserRolesUsecaseImpl struct {
	userRoleRepository repository.UserRoleRepository
	helper             utils.Helper
}

func NewListUserRolesUsecase(userRoleRepository repository.UserRoleRepository, helper utils.Helper) ListUserRolesUsecase {
	return &ListUserRolesUsecaseImpl{
		userRoleRepository: userRoleRepository,
		helper:             helper,
	}
}

func (u *ListUserRolesUsecaseImpl) Execute(
	ctx context.Context,
	pagination common.Pagination,
	filter entity.UserRoleFilter,
) (common.PaginationResult[*entity.UserRole], error) {
	userRoles, total, err := u.userRoleRepository.List(ctx, pagination, filter)
	if err != nil {
		return common.PaginationResult[*entity.UserRole]{}, ErrListUserRoles
	}
	return common.PaginationResult[*entity.UserRole]{
		Data:       userRoles,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(pagination.PageSize)),
	}, nil
}
