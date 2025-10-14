package repo

import (
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type roleRepository struct {
	db pg.DBI
}

func NewRoleRepository(db *pg.DB) repository.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (rr *roleRepository) CreateRole(role entity.Role) error {
	_, err := rr.db.Model(&role).Insert()
	return err
}

func (rr *roleRepository) GetRoleByID(id string) (entity.Role, error) {
	var role entity.Role
	err := rr.db.Model(&role).Where("id = ?", id).Select()
	return role, err
}

func (rr *roleRepository) GetRoleByName(name string) (entity.Role, error) {
	var role entity.Role
	err := rr.db.Model(&role).Where("name = ?", name).Select()
	return role, err
}

func (rr *roleRepository) GetAllRoles() ([]entity.Role, error) {
	var roles []entity.Role
	err := rr.db.Model(&roles).Select()
	return roles, err
}

func (rr *roleRepository) UpdateRole(id string, role entity.Role) (entity.Role, error) {
	_, err := rr.db.Model(&role).Where("id = ?", id).UpdateNotZero()
	return role, err
}

func (rr *roleRepository) DeleteByID(id string) error {
	var role entity.Role
	_, err := rr.db.Model(&role).Where("id = ?", id).Delete()
	return err
}

func (rr *roleRepository) CheckRoleExist(name string) (bool, error) {
	var role entity.Role
	count, err := rr.db.Model(&role).Where("name = ?", name).Count()
	isExist := count > 0
	return isExist, err
}
