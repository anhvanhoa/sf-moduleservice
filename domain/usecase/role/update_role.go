package role

import (
	"module-service/domain/entity"
	"module-service/domain/repository"
	"time"
)

type UpdateRoleUsecase interface {
	Excute(id string, role entity.Role) (entity.Role, error)
}

type updateRoleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewUpdateRoleUsecase(roleRepo repository.RoleRepository) UpdateRoleUsecase {
	return &updateRoleUsecase{
		roleRepo: roleRepo,
	}
}

func (u *updateRoleUsecase) Excute(id string, role entity.Role) (entity.Role, error) {
	now := time.Now()
	role.UpdatedAt = &now
	updatedRole, err := u.roleRepo.UpdateRole(id, role)
	if err != nil {
		return entity.Role{}, ErrUpdateRole
	}
	return updatedRole, nil
}
