package repository

import (
	"context"

	"module-service/domain/common"
	"module-service/domain/entity"
)

// ModuleChildRepository defines the interface for ModuleChild entity operations
type ModuleChildRepository interface {
	// Create operations
	Create(ctx context.Context, moduleChild *entity.ModuleChild) error
	CreateMany(ctx context.Context, moduleChildren []*entity.ModuleChild) error

	// Read operations
	GetByID(ctx context.Context, id string) (*entity.ModuleChild, error)
	GetByModuleID(ctx context.Context, moduleID string) ([]*entity.ModuleChild, error)
	GetByPath(ctx context.Context, path string) (*entity.ModuleChild, error)
	GetByPathAndMethod(ctx context.Context, path, method string) (*entity.ModuleChild, error)
	GetByStatus(ctx context.Context, status common.Status) ([]*entity.ModuleChild, error)
	GetByIsPrivate(ctx context.Context, isPrivate bool) ([]*entity.ModuleChild, error)
	GetAll(ctx context.Context) ([]*entity.ModuleChild, error)
	GetWithPagination(ctx context.Context, pagination *common.Pagination) ([]*entity.ModuleChild, int64, error)
	ExistsByID(ctx context.Context, id string) (bool, error)
	ExistsByPathAndMethod(ctx context.Context, path, method string) (bool, error)

	// Update operations
	Update(ctx context.Context, moduleChild *entity.ModuleChild) error
	UpdateStatus(ctx context.Context, id string, status common.Status) error
	UpdateName(ctx context.Context, id string, name string) error
	UpdatePath(ctx context.Context, id string, path string) error
	UpdateMethod(ctx context.Context, id string, method string) error
	UpdateIsPrivate(ctx context.Context, id string, isPrivate bool) error

	// Delete operations
	DeleteByID(ctx context.Context, id string) error
	DeleteByModuleID(ctx context.Context, moduleID string) error
	DeleteByStatus(ctx context.Context, status common.Status) error
	SoftDeleteByID(ctx context.Context, id string) error

	// Search operations
	SearchByName(ctx context.Context, name string) ([]*entity.ModuleChild, error)
	SearchByPath(ctx context.Context, path string) ([]*entity.ModuleChild, error)
	SearchByMethod(ctx context.Context, method string) ([]*entity.ModuleChild, error)

	// Relationship operations
	GetWithModule(ctx context.Context, id string) (*entity.ModuleChild, *entity.Module, error)
	GetModuleChildrenWithModule(ctx context.Context, moduleID string) ([]*entity.ModuleChild, *entity.Module, error)
}
