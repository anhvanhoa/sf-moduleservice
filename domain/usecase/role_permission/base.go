package role_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type RolePermissionUsecaseI interface {
	Create(ctx context.Context, rolePermission *entity.RolePermission) error
	List(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) (common.PaginationResult[*entity.RolePermission], error)
	Delete(ctx context.Context, roleID, permissionID string) error
	DeleteByPermissionID(ctx context.Context, permissionID string) error
	Count(ctx context.Context) (int64, error)
	CountByRoleID(ctx context.Context, roleID string) (int64, error)
	CountByPermissionID(ctx context.Context, permissionID string) (int64, error)
	Exists(ctx context.Context, roleID, permissionID string) (bool, error)
}

type RolePermissionUsecaseImpl struct {
	createRolePermissionUsecase CreateRolePermissionUsecase
	listRolePermissionsUsecase  ListRolePermissionsUsecase
	deleteRolePermissionUsecase DeleteRolePermissionUsecase
	deleteByPermissionIDUsecase DeleteByPermissionIDUsecase
	countRolePermissionsUsecase CountRolePermissionsUsecase
	countByRoleIDUsecase        CountByRoleIDUsecase
	countByPermissionIDUsecase  CountByPermissionIDUsecase
	existsRolePermissionUsecase ExistsRolePermissionUsecase
}

func NewRolePermissionUsecase(rolePermissionRepository repository.RolePermissionRepository, helper utils.Helper) RolePermissionUsecaseI {
	return &RolePermissionUsecaseImpl{
		createRolePermissionUsecase: NewCreateRolePermissionUsecase(rolePermissionRepository),
		listRolePermissionsUsecase:  NewListRolePermissionsUsecase(rolePermissionRepository, helper),
		deleteRolePermissionUsecase: NewDeleteRolePermissionUsecase(rolePermissionRepository),
		deleteByPermissionIDUsecase: NewDeleteByPermissionIDUsecase(rolePermissionRepository),
		countRolePermissionsUsecase: NewCountRolePermissionsUsecase(rolePermissionRepository),
		countByRoleIDUsecase:        NewCountByRoleIDUsecase(rolePermissionRepository),
		countByPermissionIDUsecase:  NewCountByPermissionIDUsecase(rolePermissionRepository),
		existsRolePermissionUsecase: NewExistsRolePermissionUsecase(rolePermissionRepository),
	}
}

func (u *RolePermissionUsecaseImpl) Create(ctx context.Context, rolePermission *entity.RolePermission) error {
	return u.createRolePermissionUsecase.Execute(ctx, rolePermission)
}

func (u *RolePermissionUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) (common.PaginationResult[*entity.RolePermission], error) {
	return u.listRolePermissionsUsecase.Execute(ctx, pagination, filter)
}

func (u *RolePermissionUsecaseImpl) Delete(ctx context.Context, roleID, permissionID string) error {
	return u.deleteRolePermissionUsecase.Execute(ctx, roleID, permissionID)
}

func (u *RolePermissionUsecaseImpl) DeleteByPermissionID(ctx context.Context, permissionID string) error {
	return u.deleteByPermissionIDUsecase.Execute(ctx, permissionID)
}

func (u *RolePermissionUsecaseImpl) Count(ctx context.Context) (int64, error) {
	return u.countRolePermissionsUsecase.Execute(ctx)
}

func (u *RolePermissionUsecaseImpl) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	return u.countByRoleIDUsecase.Execute(ctx, roleID)
}

func (u *RolePermissionUsecaseImpl) CountByPermissionID(ctx context.Context, permissionID string) (int64, error) {
	return u.countByPermissionIDUsecase.Execute(ctx, permissionID)
}

func (u *RolePermissionUsecaseImpl) Exists(ctx context.Context, roleID, permissionID string) (bool, error) {
	return u.existsRolePermissionUsecase.Execute(ctx, roleID, permissionID)
}
