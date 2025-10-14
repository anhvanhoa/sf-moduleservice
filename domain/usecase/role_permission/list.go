package role_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListRolePermissionsUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) (common.PaginationResult[*entity.RolePermission], error)
}

type ListRolePermissionsUsecaseImpl struct {
	rolePermissionRepository repository.RolePermissionRepository
	helper                   utils.Helper
}

func NewListRolePermissionsUsecase(
	rolePermissionRepository repository.RolePermissionRepository,
	helper utils.Helper,
) ListRolePermissionsUsecase {
	return &ListRolePermissionsUsecaseImpl{
		rolePermissionRepository: rolePermissionRepository,
		helper:                   helper,
	}
}

func (u *ListRolePermissionsUsecaseImpl) Execute(
	ctx context.Context,
	pagination common.Pagination,
	filter entity.RolePermissionFilter,
) (common.PaginationResult[*entity.RolePermission], error) {
	rolePermissions, total, err := u.rolePermissionRepository.List(ctx, pagination, filter)
	if err != nil {
		return common.PaginationResult[*entity.RolePermission]{}, ErrListRolePermissions
	}
	return common.PaginationResult[*entity.RolePermission]{
		Data:       rolePermissions,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(pagination.PageSize)),
	}, nil
}
