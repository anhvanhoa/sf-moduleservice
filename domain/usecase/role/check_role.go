package role

import (
	"module-service/domain/repository"
)

type CheckRoleUsecase interface {
	Excute(name string) (bool, error)
}

type checkRoleUsecase struct {
	roleRepo repository.RoleRepository
}

func NewCheckRoleUsecase(roleRepo repository.RoleRepository) CheckRoleUsecase {
	return &checkRoleUsecase{
		roleRepo: roleRepo,
	}
}

func (c *checkRoleUsecase) Excute(name string) (bool, error) {
	isExist, err := c.roleRepo.CheckRoleExist(name)
	if err != nil {
		return false, ErrCheckRole
	}
	return isExist, nil
}
