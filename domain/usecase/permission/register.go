package permission

import (
	"context"
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

func (u *RegisterPermissionUsecaseImpl) Execute(ctx context.Context, permissions []*entity.Permission) error {
	existingPermissions, _, err := u.permissionRepository.List(ctx, nil, nil)
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
	diffPermissions := make([]*entity.Permission, 0)
	for _, permission := range permissions {
		found := false
		for _, existingPermission := range existingPermissions {
			if existingPermission.Resource == permission.Resource && existingPermission.Action == permission.Action {
				found = true
				break
			}
		}
		if !found {
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

	// Set new permissions to cache
	for _, permission := range diffPermissions {
		cacheKey := permission.Resource + "." + permission.Action
		u.cacher.Set(cacheKey, []byte("false"), 0)
	}

	return nil
}

func (u *RegisterPermissionUsecaseImpl) RemovePermissions(ctx context.Context, permissions, existingPermissions []*entity.Permission) error {
	removePermissions := make([]*entity.Permission, 0)
	for _, existingPermission := range existingPermissions {
		found := false
		for _, permission := range permissions {
			if existingPermission.Resource == permission.Resource && existingPermission.Action == permission.Action {
				found = true
				break
			}
		}
		if !found {
			removePermissions = append(removePermissions, existingPermission)
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

	// Remove permissions from cache
	for _, permission := range removePermissions {
		cacheKey := permission.Resource + "." + permission.Action
		u.cacher.Delete(cacheKey)
	}

	return nil
}

func (u *RegisterPermissionUsecaseImpl) UpdateExistingPermissionsCache(ctx context.Context, permissions, existingPermissions []*entity.Permission) error {
	// Update cache for existing permissions that are still valid
	for _, permission := range permissions {
		for _, existingPermission := range existingPermissions {
			if existingPermission.Resource == permission.Resource && existingPermission.Action == permission.Action {
				// This permission exists in DB and is still valid, update cache
				cacheKey := permission.Resource + "." + permission.Action
				u.cacher.Set(cacheKey, []byte("false"), 0)
				break
			}
		}
	}
	return nil
}
