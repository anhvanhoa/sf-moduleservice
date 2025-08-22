package repo

import (
	"context"

	"module-service/domain/common"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type moduleRepository struct {
	db *pg.DB
}

// NewModuleRepository creates a new instance of ModuleRepository
func NewModuleRepository(db *pg.DB) repository.ModuleRepository {
	return &moduleRepository{
		db: db,
	}
}

// Create operations
func (r *moduleRepository) Create(ctx context.Context, module *entity.Module) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(module).Insert()
	return err
}

func (r *moduleRepository) CreateMany(ctx context.Context, modules []*entity.Module) error {
	if len(modules) == 0 {
		return nil
	}
	db := getTx(ctx, r.db)
	_, err := db.Model(&modules).Insert()
	return err
}

// Read operations
func (r *moduleRepository) GetByID(ctx context.Context, id string) (*entity.Module, error) {
	var module entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&module).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &module, nil
}

func (r *moduleRepository) GetByName(ctx context.Context, name string) (*entity.Module, error) {
	var module entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&module).Where("name = ?", name).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &module, nil
}

func (r *moduleRepository) GetByStatus(ctx context.Context, status common.Status) ([]*entity.Module, error) {
	var modules []*entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&modules).Where("status = ?", status).Select()
	return modules, err
}

func (r *moduleRepository) GetAll(ctx context.Context) ([]*entity.Module, error) {
	var modules []*entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&modules).Select()
	return modules, err
}

func (r *moduleRepository) GetWithPagination(ctx context.Context, pagination *common.Pagination) ([]*entity.Module, int64, error) {
	var modules []*entity.Module
	var total int64

	db := getTx(ctx, r.db)

	// Get total count
	count, err := db.Model(&modules).Count()
	if err != nil {
		return nil, 0, err
	}
	total = int64(count)

	// Get paginated results
	query := db.Model(&modules)
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
	return modules, total, err
}

func (r *moduleRepository) ExistsByID(ctx context.Context, id string) (bool, error) {
	db := getTx(ctx, r.db)
	exists, err := db.Model((*entity.Module)(nil)).Where("id = ?", id).Exists()
	return exists, err
}

func (r *moduleRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	db := getTx(ctx, r.db)
	exists, err := db.Model((*entity.Module)(nil)).Where("name = ?", name).Exists()
	return exists, err
}

// Update operations
func (r *moduleRepository) Update(ctx context.Context, module *entity.Module) error {
	db := getTx(ctx, r.db)
	_, err := db.Model(module).WherePK().Update()
	return err
}

func (r *moduleRepository) UpdateStatus(ctx context.Context, id string, status common.Status) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).
		Set("status = ?", status).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleRepository) UpdateName(ctx context.Context, id string, name string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).
		Set("name = ?", name).
		Where("id = ?", id).
		Update()
	return err
}

func (r *moduleRepository) UpdateDescription(ctx context.Context, id string, description string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).
		Set("description = ?", description).
		Where("id = ?", id).
		Update()
	return err
}

// Delete operations
func (r *moduleRepository) DeleteByID(ctx context.Context, id string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).Where("id = ?", id).Delete()
	return err
}

func (r *moduleRepository) DeleteByStatus(ctx context.Context, status common.Status) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).Where("status = ?", status).Delete()
	return err
}

func (r *moduleRepository) SoftDeleteByID(ctx context.Context, id string) error {
	db := getTx(ctx, r.db)
	_, err := db.Model((*entity.Module)(nil)).
		Set("status = ?", common.StatusInactive).
		Where("id = ?", id).
		Update()
	return err
}

// Search operations
func (r *moduleRepository) SearchByName(ctx context.Context, name string) ([]*entity.Module, error) {
	var modules []*entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&modules).Where("name ILIKE ?", "%"+name+"%").Select()
	return modules, err
}

func (r *moduleRepository) SearchByDescription(ctx context.Context, description string) ([]*entity.Module, error) {
	var modules []*entity.Module
	db := getTx(ctx, r.db)
	err := db.Model(&modules).Where("description ILIKE ?", "%"+description+"%").Select()
	return modules, err
}
