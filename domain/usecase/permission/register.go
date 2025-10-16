package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type RegisterPermissionsUsecase interface {
	Execute(ctx context.Context, permissions []*entity.Permission) error
	DiffPermissions(ctx context.Context, permissions []*entity.Permission, existingPermissions []*entity.Permission) error
	RemovePermissions(ctx context.Context, permissions, existingPermissions []*entity.Permission) error
}

type RegisterPermissionUsecaseImpl struct {
	permissionRepository repository.PermissionRepository
}

func NewRegisterPermissionsUsecase(permissionRepository repository.PermissionRepository) RegisterPermissionsUsecase {
	return &RegisterPermissionUsecaseImpl{
		permissionRepository: permissionRepository,
	}
}

func (u *RegisterPermissionUsecaseImpl) Execute(ctx context.Context, permissions []*entity.Permission) error {
	existingPermissions, _, err := u.permissionRepository.List(ctx, nil, nil)
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
	return u.permissionRepository.CreateMany(ctx, diffPermissions)
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
	return u.permissionRepository.DeleteMany(ctx, removePermissions)
}
