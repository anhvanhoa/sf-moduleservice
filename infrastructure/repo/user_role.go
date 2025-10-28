package repo

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/domain/user_context"
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

func (r *userRoleRepository) GetUserPermissions(ctx context.Context, userID string) (user_context.UserContext, error) {
	uCtx := user_context.NewUserContext()
	uCtx.UserID = userID

	roleMap := make(map[string]bool, 2)
	permissionMap := make(map[string]user_context.Permission, 20)
	scopes := make([]user_context.Scope, 0, 10)

	var rolePermissionResults []struct {
		RoleName string
		Resource string
		Action   string
	}

	rolePermissionQuery := `
		SELECT DISTINCT
			r.name as role_name,
			p.resource,
			p.action
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		INNER JOIN role_permissions rp ON r.id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND r.status = 'active'
		ORDER BY r.name, p.resource, p.action
	`

	_, err := r.db.Query(&rolePermissionResults, rolePermissionQuery, userID)
	if err != nil {
		return *uCtx, err
	}

	for _, result := range rolePermissionResults {
		roleMap[result.RoleName] = true
		permissionKey := result.Resource + "." + result.Action
		permissionMap[permissionKey] = user_context.Permission{
			Resource: result.Resource,
			Action:   result.Action,
		}
	}

	var scopeResults []struct {
		ResourceType string
		ResourceData map[string]string
		Action       string
	}

	scopeQuery := `
		SELECT 
			resource_type,
			resource_data,
			action
		FROM resource_permissions
		WHERE user_id = ?
		ORDER BY resource_type, action
	`

	_, err = r.db.Query(&scopeResults, scopeQuery, userID)
	if err != nil {
		return *uCtx, err
	}

	for _, result := range scopeResults {
		scopes = append(scopes, user_context.Scope{
			Resource:     result.ResourceType,
			ResourceData: result.ResourceData,
			Action:       result.Action,
		})
	}

	roles := make([]string, 0, len(roleMap))
	for role := range roleMap {
		roles = append(roles, role)
	}

	permissions := make([]user_context.Permission, 0, len(permissionMap))
	for _, permission := range permissionMap {
		permissions = append(permissions, permission)
	}

	uCtx.Roles = roles
	uCtx.Permissions = permissions
	uCtx.Scopes = scopes
	return *uCtx, nil
}
