package permission

import (
	"context"
	"fmt"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/cache"
)

type RegisterPermissionsUsecase interface {
	Execute(ctx context.Context, permissions []*entity.Permission) error
	DiffPermissions(ctx context.Context, permissions []*entity.Permission, existingPermissions []*entity.Permission) error
	RemovePermissions(ctx context.Context, permissions, existingPermissions []*entity.Permission) error
	UpdateExistingPermissionsCache(ctx context.Context, permissions, existingPermissions []*entity.Permission) error
}

type RegisterPermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
	cacher               cache.CacheI
}

func NewRegisterPermissionsUsecase(permissionRepository repository.PermissionRepository, cacher cache.CacheI) RegisterPermissionsUsecase {
	return &RegisterPermissionUsecaseImpl{
		permissionRepository: permissionRepository,
		cacher:               cacher,
	}
}

// Helper methods for optimization

// getPermissionKey creates a unique key for permission lookup
func (u *RegisterPermissionUsecaseImpl) getPermissionKey(resource, action string) string {
	return fmt.Sprintf("%s.%s", resource, action)
}

// batchSetCache sets multiple permissions in cache with the same value
func (u *RegisterPermissionUsecaseImpl) batchSetCache(permissions []*entity.Permission, value string) {
	for _, permission := range permissions {
		cacheKey := u.getPermissionKey(permission.Resource, permission.Action)
		if permission.Action == "Check" {
			value = "true"
		}
		u.cacher.Set(cacheKey, []byte(value), 0)
	}
}

// batchDeleteCache removes multiple permissions from cache
func (u *RegisterPermissionUsecaseImpl) batchDeleteCache(permissions []*entity.Permission) {
	for _, permission := range permissions {
		cacheKey := u.getPermissionKey(permission.Resource, permission.Action)
		u.cacher.Delete(cacheKey)
	}
}

// batchUpdateCache updates cache for existing permissions with their IsPublic status
func (u *RegisterPermissionUsecaseImpl) batchUpdateCache(permissions []*entity.Permission) {
	for _, permission := range permissions {
		cacheKey := u.getPermissionKey(permission.Resource, permission.Action)
		isPublic := "false"
		if permission.IsPublic {
			isPublic = "true"
		}
		if permission.Action == "Check" {
			isPublic = "true"
		}
		u.cacher.Set(cacheKey, []byte(isPublic), 0)
	}
}

func (u *RegisterPermissionUsecaseImpl) Execute(ctx context.Context, permissions []*entity.Permission) error {
	filter := &entity.PermissionFilter{
		Resource: []string{},
	}
	resources := make(map[string]bool)
	for _, permission := range permissions {
		resources[permission.Resource] = true
	}
	for resource := range resources {
		filter.Resource = append(filter.Resource, resource)
	}
	existingPermissions, _, err := u.permissionRepository.List(ctx, nil, filter)
	if err != nil {
		return err
	}

	// Update cache for existing permissions that are still valid
	err = u.UpdateExistingPermissionsCache(ctx, permissions, existingPermissions)
	if err != nil {
		return err
	}

	err = u.RemovePermissions(ctx, permissions, existingPermissions)
	if err != nil {
		return err
	}
	err = u.DiffPermissions(ctx, permissions, existingPermissions)
	if err != nil {
		return err
	}
	return nil
}

func (u *RegisterPermissionUsecaseImpl) DiffPermissions(ctx context.Context, permissions, existingPermissions []*entity.Permission) error {
	// Create a map for O(1) lookup of existing permissions
	existingMap := make(map[string]*entity.Permission)
	for _, existing := range existingPermissions {
		key := u.getPermissionKey(existing.Resource, existing.Action)
		existingMap[key] = existing
	}

	// Find permissions that don't exist
	diffPermissions := make([]*entity.Permission, 0, len(permissions))
	for _, permission := range permissions {
		key := u.getPermissionKey(permission.Resource, permission.Action)
		if _, exists := existingMap[key]; !exists {
			if permission.Action == "Check" {
				permission.IsPublic = true
			}
			diffPermissions = append(diffPermissions, permission)
		}
	}

	if len(diffPermissions) == 0 {
		return nil
	}

	// Create new permissions in database
	err := u.permissionRepository.CreateMany(ctx, diffPermissions)
	if err != nil {
		return err
	}

	// Batch cache operations for new permissions
	u.batchSetCache(diffPermissions, "false")

	return nil
}

func (u *RegisterPermissionUsecaseImpl) RemovePermissions(ctx context.Context, permissions, existingPermissions []*entity.Permission) error {
	// Create a map for O(1) lookup of current permissions
	currentMap := make(map[string]bool)
	for _, permission := range permissions {
		key := u.getPermissionKey(permission.Resource, permission.Action)
		currentMap[key] = true
	}

	// Find permissions to remove
	removePermissions := make([]*entity.Permission, 0, len(existingPermissions))
	for _, existing := range existingPermissions {
		key := u.getPermissionKey(existing.Resource, existing.Action)
		if !currentMap[key] {
			removePermissions = append(removePermissions, existing)
		}
	}

	if len(removePermissions) == 0 {
		return nil
	}

	// Delete permissions from database
	err := u.permissionRepository.DeleteMany(ctx, removePermissions)
	if err != nil {
		return err
	}

	// Batch remove permissions from cache
	u.batchDeleteCache(removePermissions)

	return nil
}

func (u *RegisterPermissionUsecaseImpl) UpdateExistingPermissionsCache(ctx context.Context, permissions, existingPermissions []*entity.Permission) error {
	// Create a map for O(1) lookup of existing permissions
	existingMap := make(map[string]*entity.Permission)
	for _, existing := range existingPermissions {
		key := u.getPermissionKey(existing.Resource, existing.Action)
		existingMap[key] = existing
	}

	// Update cache for existing permissions that are still valid
	permissionsToUpdate := make([]*entity.Permission, 0, len(permissions))
	for _, permission := range permissions {
		key := u.getPermissionKey(permission.Resource, permission.Action)
		if existing, found := existingMap[key]; found {
			permissionsToUpdate = append(permissionsToUpdate, existing)
		}
	}

	// Batch update cache
	u.batchUpdateCache(permissionsToUpdate)
	return nil
}
