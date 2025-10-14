package repo

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type userRoleRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewUserRoleRepository(db *pg.DB, helper utils.Helper) repository.UserRoleRepository {
	return &userRoleRepository{
		db:     db,
		helper: helper,
	}
}

func (r *userRoleRepository) Create(ctx context.Context, userRole *entity.UserRole) error {
	_, err := r.db.Model(userRole).Context(ctx).Insert()
	return err
}

func (r *userRoleRepository) List(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error) {
	var userRoles []*entity.UserRole
	query := r.db.Model(&userRoles).Context(ctx)

	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.RoleID != "" {
		query = query.Where("role_id = ?", filter.RoleID)
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

	return userRoles, int64(total), nil
}

func (r *userRoleRepository) Delete(ctx context.Context, userID, roleID string) error {
	_, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).Delete()
	return err
}

func (r *userRoleRepository) DeleteByUserID(ctx context.Context, userID string) error {
	_, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("user_id = ?", userID).Delete()
	return err
}

func (r *userRoleRepository) DeleteByRoleID(ctx context.Context, roleID string) error {
	_, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("role_id = ?", roleID).Delete()
	return err
}

func (r *userRoleRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Model(&entity.UserRole{}).Context(ctx).Count()
	return int64(count), err
}

func (r *userRoleRepository) CountByUserID(ctx context.Context, userID string) (int64, error) {
	count, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("user_id = ?", userID).Count()
	return int64(count), err
}

func (r *userRoleRepository) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	count, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("role_id = ?", roleID).Count()
	return int64(count), err
}

func (r *userRoleRepository) Exists(ctx context.Context, userID, roleID string) (bool, error) {
	count, err := r.db.Model(&entity.UserRole{}).Context(ctx).
		Where("user_id = ? AND role_id = ?", userID, roleID).Count()
	return count > 0, err
}
