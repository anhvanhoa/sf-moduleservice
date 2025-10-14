package repository

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type UserRoleRepository interface {
	Create(ctx context.Context, userRole *entity.UserRole) error
	List(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error)
	Delete(ctx context.Context, userID, roleID string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteByRoleID(ctx context.Context, roleID string) error
	Count(ctx context.Context) (int64, error)
	CountByUserID(ctx context.Context, userID string) (int64, error)
	CountByRoleID(ctx context.Context, roleID string) (int64, error)
	Exists(ctx context.Context, userID, roleID string) (bool, error)
}
