package repository

import (
	"context"
	"module-service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission *entity.Permission) error
	GetByID(ctx context.Context, id string) (*entity.Permission, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) ([]*entity.Permission, int64, error)
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id string) error
	DeleteByResourceAndAction(ctx context.Context, resource, action string) error
	CountByResource(ctx context.Context, resource string) (int64, error)
}
