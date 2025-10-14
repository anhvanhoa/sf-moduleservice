package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListResourcePermissionsUsecase interface {
	Execute(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) (common.PaginationResult[*entity.ResourcePermission], error)
}

type ListResourcePermissionsUsecaseImpl struct {
	resourcePermissionRepository repository.ResourcePermissionRepository
	helper                       utils.Helper
}

func NewListResourcePermissionsUsecase(
	resourcePermissionRepository repository.ResourcePermissionRepository,
	helper utils.Helper,
) ListResourcePermissionsUsecase {
	return &ListResourcePermissionsUsecaseImpl{
		resourcePermissionRepository,
		helper,
	}
}

func (u *ListResourcePermissionsUsecaseImpl) Execute(
	ctx context.Context, pagination common.Pagination,
	filter entity.ResourcePermissionFilter,
) (common.PaginationResult[*entity.ResourcePermission], error) {
	resourcePermissions, total, err := u.resourcePermissionRepository.List(ctx, pagination, filter)
	if err != nil {
		return common.PaginationResult[*entity.ResourcePermission]{}, ErrListResourcePermissions
	}
	return common.PaginationResult[*entity.ResourcePermission]{
		Data:       resourcePermissions,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(pagination.PageSize)),
	}, nil
}
