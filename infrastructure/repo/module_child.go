package repo

import (
	"context"

	"module-service/domain/common"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type moduleChildRepository struct {
	db *pg.DB
}

// NewModuleChildRepository creates a new instance of ModuleChildRepository
func NewModuleChildRepository(db *pg.DB) repository.ModuleChildRepository {
	return &moduleChildRepository{
		db: db,
	}
}

// Create operations
func (r *moduleChildRepository) Create(ctx context.Context, moduleChild *entity.ModuleChild) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(moduleChild).Insert()
	return err
}

func (r *moduleChildRepository) CreateMany(ctx context.Context, moduleChildren []*entity.ModuleChild) error {
	if len(moduleChildren) == 0 {
		return nil
	}
	db := getTx(ctx, r.db)
	_, err := db.Model(&moduleChildren).Insert()
	return err
}

// Read operations
func (r *moduleChildRepository) GetByID(ctx context.Context, id string) (*entity.ModuleChild, error) {
	var moduleChild entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChild).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &moduleChild, nil
}

func (r *moduleChildRepository) GetByModuleID(ctx context.Context, moduleID string) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("module_id = ?", moduleID).Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) GetByPath(ctx context.Context, path string) (*entity.ModuleChild, error) {
	var moduleChild entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChild).Where("path = ?", path).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &moduleChild, nil
}

func (r *moduleChildRepository) GetByPathAndMethod(ctx context.Context, path, method string) (*entity.ModuleChild, error) {
	var moduleChild entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChild).Where("path = ? AND method = ?", path, method).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &moduleChild, nil
}

func (r *moduleChildRepository) GetByStatus(ctx context.Context, status common.Status) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("status = ?", status).Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) GetByIsPrivate(ctx context.Context, isPrivate bool) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("is_private = ?", isPrivate).Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) GetAll(ctx context.Context) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) GetWithPagination(ctx context.Context, pagination *common.Pagination) ([]*entity.ModuleChild, int64, error) {
	var moduleChildren []*entity.ModuleChild
	var total int64

	db := getTx(ctx, r.db)

	// Get total count
	count, err := db.Model(&moduleChildren).Count()
	if err != nil {
		return nil, 0, err
	}
	total = int64(count)

	// Get paginated results
	query := db.Model(&moduleChildren)
	if pagination != nil {
		if pagination.PageSize > 0 {
			query = query.Limit(pagination.PageSize)
		}
		if pagination.Page > 1 {
			offset := (pagination.Page - 1) * pagination.PageSize
			query = query.Offset(offset)
		}
	}

	err = query.Select()
	return moduleChildren, total, err
}

func (r *moduleChildRepository) ExistsByID(ctx context.Context, id string) (bool, error) {
	db := getTx(ctx, r.db)
	exists, err := db.Model((*entity.ModuleChild)(nil)).Where("id = ?", id).Exists()
	return exists, err
}

func (r *moduleChildRepository) ExistsByPathAndMethod(ctx context.Context, path, method string) (bool, error) {
	db := getTx(ctx, r.db)
	exists, err := db.Model((*entity.ModuleChild)(nil)).Where("path = ? AND method = ?", path, method).Exists()
	return exists, err
}

// Update operations
func (r *moduleChildRepository) Update(ctx context.Context, moduleChild *entity.ModuleChild) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(moduleChild).WherePK().Update()
	return err
}

func (r *moduleChildRepository) UpdateStatus(ctx context.Context, id string, status common.Status) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("status = ?", status).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleChildRepository) UpdateName(ctx context.Context, id string, name string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("name = ?", name).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleChildRepository) UpdatePath(ctx context.Context, id string, path string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("path = ?", path).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleChildRepository) UpdateMethod(ctx context.Context, id string, method string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("method = ?", method).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleChildRepository) UpdateIsPrivate(ctx context.Context, id string, isPrivate bool) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("is_private = ?", isPrivate).
		Where("id = ?", id).
		Update()
	return err
}

// Delete operations
func (r *moduleChildRepository) DeleteByID(ctx context.Context, id string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).Where("id = ?", id).Delete()
	return err
}

func (r *moduleChildRepository) DeleteByModuleID(ctx context.Context, moduleID string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).Where("module_id = ?", moduleID).Delete()
	return err
}

func (r *moduleChildRepository) DeleteByStatus(ctx context.Context, status common.Status) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).Where("status = ?", status).Delete()
	return err
}

func (r *moduleChildRepository) SoftDeleteByID(ctx context.Context, id string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.ModuleChild)(nil)).
		Set("status = ?", common.StatusInactive).
		Where("id = ?", id).
		Update()
	return err
}

// Search operations
func (r *moduleChildRepository) SearchByName(ctx context.Context, name string) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("name ILIKE ?", "%"+name+"%").Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) SearchByPath(ctx context.Context, path string) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("path ILIKE ?", "%"+path+"%").Select()
	return moduleChildren, err
}

func (r *moduleChildRepository) SearchByMethod(ctx context.Context, method string) ([]*entity.ModuleChild, error) {
	var moduleChildren []*entity.ModuleChild
	db := getTx(ctx, r.db)
	err := db.Model(&moduleChildren).Where("method ILIKE ?", "%"+method+"%").Select()
	return moduleChildren, err
}

// Relationship operations
func (r *moduleChildRepository) GetWithModule(ctx context.Context, id string) (*entity.ModuleChild, *entity.Module, error) {
	var moduleChild entity.ModuleChild
	var module entity.Module

	db := getTx(ctx, r.db)

	// Get module child with its module
	err := db.Model(&moduleChild).
		Join("JOIN modules m ON m.id = mc.module_id").
		Where("mc.id = ?", id).
		Select(&moduleChild, &module)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return &moduleChild, &module, nil
}

func (r *moduleChildRepository) GetModuleChildrenWithModule(ctx context.Context, moduleID string) ([]*entity.ModuleChild, *entity.Module, error) {
	var moduleChildren []*entity.ModuleChild
	var module entity.Module

	db := getTx(ctx, r.db)

	// Get module first
	err := db.Model(&module).Where("id = ?", moduleID).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	// Get module children
	err = db.Model(&moduleChildren).Where("module_id = ?", moduleID).Select()
	if err != nil {
		return nil, nil, err
	}

	return moduleChildren, &module, nil
}
