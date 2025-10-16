package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListPermissionsUsecase interface {
	Execute(ctx context.Context, pagination *common.Pagination, filter *entity.PermissionFilter) (common.PaginationResult[*entity.Permission], error)
}

type ListPermissionsUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
	helper               utils.Helper
}

func NewListPermissionsUsecase(permissionRepository repository.PermissionRepository, helper utils.Helper) ListPermissionsUsecase {
	return &ListPermissionsUsecaseImpl{
		permissionRepository: permissionRepository,
		helper:               helper,
	}
}

func (u *ListPermissionsUsecaseImpl) Execute(
	ctx context.Context, pagination *common.Pagination, filter *entity.PermissionFilter,
) (common.PaginationResult[*entity.Permission], error) {
	permissions, total, err := u.permissionRepository.List(ctx, pagination, filter)
	if err != nil {
		return common.PaginationResult[*entity.Permission]{}, ErrListPermissions
	}
	return common.PaginationResult[*entity.Permission]{
		Data:       permissions,
		Total:      total,
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalPages: u.helper.CalculateTotalPages(total, int64(pagination.PageSize)),
	}, nil
}
