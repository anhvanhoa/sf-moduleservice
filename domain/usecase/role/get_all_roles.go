package role

import (
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetAllRolesUsecase interface {
	Excute() ([]entity.Role, error)
}

type getAllRolesUsecase struct {
	roleRepo repository.RoleRepository
}

func NewGetAllRolesUsecase(roleRepo repository.RoleRepository) GetAllRolesUsecase {
	return &getAllRolesUsecase{
		roleRepo: roleRepo,
	}
}

func (g *getAllRolesUsecase) Excute() ([]entity.Role, error) {
	roles, err := g.roleRepo.GetAllRoles()
	if err != nil {
		return nil, ErrListRoles
	}
	return roles, nil
}
