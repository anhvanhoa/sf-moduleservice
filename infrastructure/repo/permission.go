package repo

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type permissionRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewPermissionRepository(db *pg.DB, helper utils.Helper) repository.PermissionRepository {
	return &permissionRepository{
		db:     db,
		helper: helper,
	}
}

func (r *permissionRepository) Create(ctx context.Context, permission *entity.Permission) error {
	_, err := r.db.Model(permission).Context(ctx).Insert()
	return err
}

func (r *permissionRepository) CreateMany(ctx context.Context, permissions []*entity.Permission) error {
	_, err := r.db.Model(permissions).Context(ctx).Insert()
	return err
}

func (r *permissionRepository) GetByID(ctx context.Context, id string) (*entity.Permission, error) {
	permission := &entity.Permission{}
	err := r.db.Model(permission).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (r *permissionRepository) List(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) ([]*entity.Permission, int64, error) {
	var permissions []*entity.Permission
	query := r.db.Model(&permissions).Context(ctx)

	if filter.Resource != "" {
		query = query.Where("resource = ?", filter.Resource)
	}
	if filter.Action != "" {
		query = query.Where("action = ?", filter.Action)
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination.PageSize >= 0 {
		pagination.PageSize = 10
	}
	if pagination.Page >= 0 {
		pagination.Page = 1
	}
	offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
	query = query.Offset(int(offset)).Limit(int(pagination.PageSize))
	err = query.Select()
	if err != nil {
		return nil, 0, err
	}
	return permissions, int64(total), nil
}

func (r *permissionRepository) Update(ctx context.Context, permission *entity.Permission) error {
	_, err := r.db.Model(permission).Context(ctx).Where("id = ?", permission.ID).UpdateNotZero()
	return err
}

func (r *permissionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.Permission{}).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *permissionRepository) DeleteByResourceAndAction(ctx context.Context, resource, action string) error {
	_, err := r.db.Model(&entity.Permission{}).Context(ctx).
		Where("resource = ? AND action = ?", resource, action).Delete()
	return err
}

func (r *permissionRepository) CountByResource(ctx context.Context, resource string) (int64, error) {
	count, err := r.db.Model(&entity.Permission{}).Context(ctx).
		Where("resource = ?", resource).Count()
	return int64(count), err
}
