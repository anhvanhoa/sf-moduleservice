package repo

import (
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type Repositories interface {
	PermissionRepository() repository.PermissionRepository
	RoleRepository() repository.RoleRepository
	UserRoleRepository() repository.UserRoleRepository
	RolePermissionRepository() repository.RolePermissionRepository
	ResourcePermissionRepository() repository.ResourcePermissionRepository
}

type repos struct {
	permissionRepository         repository.PermissionRepository
	roleRepository               repository.RoleRepository
	userRoleRepository           repository.UserRoleRepository
	rolePermissionRepository     repository.RolePermissionRepository
	resourcePermissionRepository repository.ResourcePermissionRepository
}

func NewRepositories(db *pg.DB, helper utils.Helper) Repositories {
	return &repos{
		permissionRepository:         NewPermissionRepository(db, helper),
		roleRepository:               NewRoleRepository(db),
		userRoleRepository:           NewUserRoleRepository(db, helper),
		rolePermissionRepository:     NewRolePermissionRepository(db, helper),
		resourcePermissionRepository: NewResourcePermissionRepository(db, helper),
	}
}

func (r *repos) PermissionRepository() repository.PermissionRepository {
	return r.permissionRepository
}

func (r *repos) RoleRepository() repository.RoleRepository {
	return r.roleRepository
}

func (r *repos) UserRoleRepository() repository.UserRoleRepository {
	return r.userRoleRepository
}

func (r *repos) RolePermissionRepository() repository.RolePermissionRepository {
	return r.rolePermissionRepository
}

func (r *repos) ResourcePermissionRepository() repository.ResourcePermissionRepository {
	return r.resourcePermissionRepository
}
