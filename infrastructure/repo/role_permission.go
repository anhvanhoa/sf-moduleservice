package repo

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type rolePermissionRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewRolePermissionRepository(db *pg.DB, helper utils.Helper) repository.RolePermissionRepository {
	return &rolePermissionRepository{
		db:     db,
		helper: helper,
	}
}

func (r *rolePermissionRepository) Create(ctx context.Context, rolePermission *entity.RolePermission) error {
	_, err := r.db.Model(rolePermission).Context(ctx).Insert()
	return err
}

func (r *rolePermissionRepository) List(ctx context.Context, pagination common.Pagination, filter entity.RolePermissionFilter) ([]*entity.RolePermission, int64, error) {
	var rolePermissions []*entity.RolePermission
	query := r.db.Model(&rolePermissions).Context(ctx)

	if filter.RoleID != "" {
		query = query.Where("role_id = ?", filter.RoleID)
	}
	if filter.PermissionID != "" {
		query = query.Where("permission_id = ?", filter.PermissionID)
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
	return rolePermissions, int64(total), nil
}

func (r *rolePermissionRepository) Delete(ctx context.Context, roleID, permissionID string) error {
	_, err := r.db.Model(&entity.RolePermission{}).Context(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete()
	return err
}

func (r *rolePermissionRepository) DeleteByPermissionID(ctx context.Context, permissionID string) error {
	_, err := r.db.Model(&entity.RolePermission{}).Context(ctx).
		Where("permission_id = ?", permissionID).Delete()
	return err
}

func (r *rolePermissionRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model(&entity.RolePermission{}).Context(ctx).Count()
	return int64(count), err
}

func (r *rolePermissionRepository) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	count, err := r.db.Model(&entity.RolePermission{}).Context(ctx).
		Where("role_id = ?", roleID).Count()
	return int64(count), err
}

func (r *rolePermissionRepository) CountByPermissionID(ctx context.Context, permissionID string) (int64, error) {
	count, err := r.db.Model(&entity.RolePermission{}).Context(ctx).
		Where("permission_id = ?", permissionID).Count()
	return int64(count), err
}

func (r *rolePermissionRepository) Exists(ctx context.Context, roleID, permissionID string) (bool, error) {
	count, err := r.db.Model(&entity.RolePermission{}).Context(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).Count()
	return count > 0, err
}
