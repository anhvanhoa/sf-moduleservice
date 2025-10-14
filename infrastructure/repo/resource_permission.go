package repo

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type resourcePermissionRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewResourcePermissionRepository(db *pg.DB, helper utils.Helper) repository.ResourcePermissionRepository {
	return &resourcePermissionRepository{
		db:     db,
		helper: helper,
	}
}

func (r *resourcePermissionRepository) Create(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	_, err := r.db.Model(resourcePermission).Context(ctx).Insert()
	return err
}

func (r *resourcePermissionRepository) GetByID(ctx context.Context, id string) (*entity.ResourcePermission, error) {
	resourcePermission := &entity.ResourcePermission{}
	err := r.db.Model(resourcePermission).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return resourcePermission, nil
}

func (r *resourcePermissionRepository) List(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) ([]*entity.ResourcePermission, int64, error) {
	var resourcePermissions []*entity.ResourcePermission
	query := r.db.Model(&resourcePermissions).Context(ctx)

	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.ResourceType != "" {
		query = query.Where("resource_type = ?", filter.ResourceType)
	}
	if filter.ResourceID != "" {
		query = query.Where("resource_id = ?", filter.ResourceID)
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
	return resourcePermissions, int64(total), nil
}

func (r *resourcePermissionRepository) Update(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	_, err := r.db.Model(resourcePermission).Context(ctx).Where("id = ?", resourcePermission.ID).UpdateNotZero()
	return err
}

func (r *resourcePermissionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).Where("id = ?", id).Delete()
	return err
}

func (r *resourcePermissionRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("user_id = ?", userID).Delete()
	return err
}

func (r *resourcePermissionRepository) DeleteByResource(ctx context.Context, resourceType, resourceID string) error {
	_, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).Delete()
	return err
}

func (r *resourcePermissionRepository) DeleteByUserAndResource(ctx context.Context, userID, resourceType, resourceID string) error {
	_, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("user_id = ? AND resource_type = ? AND resource_id = ?", userID, resourceType, resourceID).Delete()
	return err
}

func (r *resourcePermissionRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).Count()
	return int64(count), err
}

func (r *resourcePermissionRepository) CountByUserID(ctx context.Context, userID string) (int64, error) {
	count, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("user_id = ?", userID).Count()
	return int64(count), err
}

func (r *resourcePermissionRepository) CountByResource(ctx context.Context, resourceType, resourceID string) (int64, error) {
	count, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("resource_type = ? AND resource_id = ?", resourceType, resourceID).Count()
	return int64(count), err
}

func (r *resourcePermissionRepository) Exists(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error) {
	count, err := r.db.Model(&entity.ResourcePermission{}).Context(ctx).
		Where("user_id = ? AND resource_type = ? AND resource_id = ? AND action = ?",
			userID, resourceType, resourceID, action).Count()
	return count > 0, err
}
