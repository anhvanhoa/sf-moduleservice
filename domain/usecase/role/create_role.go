package role

import (
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type CreateRoleUsecase interface {
	Excute(role entity.Role) error
}

type createRoleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewCreateRoleUsecase(roleRepo repository.RoleRepository) CreateRoleUsecase {
	return &createRoleUsecase{
		roleRepo: roleRepo,
	}
}

func (c *createRoleUsecase) Excute(role entity.Role) error {
	return c.roleRepo.CreateRole(role)
}
