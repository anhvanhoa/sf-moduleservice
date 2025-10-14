package repository

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type RolePermissionRepository interface {
	Create(ctx context.Context, rolePermission *entity.RolePermission) error
	List(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) ([]*entity.RolePermission, int64, error)
	Delete(ctx context.Context, roleID, permissionID string) error
	DeleteByPermissionID(ctx context.Context, permissionID string) error
	Count(ctx context.Context) (int64, error)
	CountByRoleID(ctx context.Context, roleID string) (int64, error)
	CountByPermissionID(ctx context.Context, permissionID string) (int64, error)
	Exists(ctx context.Context, roleID, permissionID string) (bool, error)
}
