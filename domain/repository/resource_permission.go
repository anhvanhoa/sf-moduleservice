package repository

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type ResourcePermissionRepository interface {
	Create(ctx context.Context, resourcePermission *entity.ResourcePermission) error
	GetByID(ctx context.Context, id string) (*entity.ResourcePermission, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) ([]*entity.ResourcePermission, int64, error)
	Update(ctx context.Context, resourcePermission *entity.ResourcePermission) error
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteByResource(ctx context.Context, resourceType, resourceID string) error
	DeleteByUserAndResource(ctx context.Context, userID, resourceType, resourceID string) error
	Count(ctx context.Context) (int64, error)
	CountByUserID(ctx context.Context, userID string) (int64, error)
	CountByResource(ctx context.Context, resourceType, resourceID string) (int64, error)
	Exists(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error)
}
