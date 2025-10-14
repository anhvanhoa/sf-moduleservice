package role

import (
	"module-service/domain/entity"
	"module-service/domain/repository"
)

type GetRoleByIDUsecase interface {
	Excute(id string) (entity.Role, error)
}

type getRoleByIDUsecase struct {
	roleRepo repository.RoleRepository
}

func NewGetRoleByIDUsecase(roleRepo repository.RoleRepository) GetRoleByIDUsecase {
	return &getRoleByIDUsecase{
		roleRepo: roleRepo,
	}
}

func (g *getRoleByIDUsecase) Excute(id string) (entity.Role, error) {
	role, err := g.roleRepo.GetRoleByID(id)
	if err != nil {
		return entity.Role{}, ErrRoleNotFound
	}
	return role, nil
}
