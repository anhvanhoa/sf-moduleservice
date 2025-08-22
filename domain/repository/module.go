package repository

import (
	"context"

	"module-service/domain/common"
	"module-service/domain/entity"
)

// ModuleRepository defines the interface for Module entity operations
type ModuleRepository interface {
	// Create operations
	Create(ctx context.Context, module *entity.Module) error
	CreateMany(ctx context.Context, modules []*entity.Module) error

	// Read operations
	GetByID(ctx context.Context, id string) (*entity.Module, error)
	GetByName(ctx context.Context, name string) (*entity.Module, error)
	GetByStatus(ctx context.Context, status common.Status) ([]*entity.Module, error)
	GetAll(ctx context.Context) ([]*entity.Module, error)
	GetWithPagination(ctx context.Context, pagination *common.Pagination) ([]*entity.Module, int64, error)
	ExistsByID(ctx context.Context, id string) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)

	// Update operations
	Update(ctx context.Context, module *entity.Module) error
	UpdateStatus(ctx context.Context, id string, status common.Status) error
	UpdateName(ctx context.Context, id string, name string) error
	UpdateDescription(ctx context.Context, id string, description string) error

	// Delete operations
	DeleteByID(ctx context.Context, id string) error
	DeleteByStatus(ctx context.Context, status common.Status) error
	SoftDeleteByID(ctx context.Context, id string) error

	// Search operations
	SearchByName(ctx context.Context, name string) ([]*entity.Module, error)
	SearchByDescription(ctx context.Context, description string) ([]*entity.Module, error)
}
