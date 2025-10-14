package user_role

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
)

type UserRoleUsecaseI interface {
	Create(ctx context.Context, userRole *entity.UserRole) error
	List(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error)
	Delete(ctx context.Context, userID, roleID string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteByRoleID(ctx context.Context, roleID string) error
	Count(ctx context.Context) (int64, error)
	CountByUserID(ctx context.Context, userID string) (int64, error)
	CountByRoleID(ctx context.Context, roleID string) (int64, error)
	Exists(ctx context.Context, userID, roleID string) (bool, error)
}

type UserRoleUsecaseImpl struct {
	createUserRoleUsecase CreateUserRoleUsecase
	listUserRolesUsecase  ListUserRolesUsecase
	deleteUserRoleUsecase DeleteUserRoleUsecase
	deleteByUserIDUsecase DeleteByUserIDUsecase
	deleteByRoleIDUsecase DeleteByRoleIDUsecase
	countUserRolesUsecase CountUserRolesUsecase
	countByUserIDUsecase  CountByUserIDUsecase
	countByRoleIDUsecase  CountByRoleIDUsecase
	existsUserRoleUsecase ExistsUserRoleUsecase
}

func NewUserRoleUsecase(userRoleRepository repository.UserRoleRepository) UserRoleUsecaseI {
	return &UserRoleUsecaseImpl{
		createUserRoleUsecase: NewCreateUserRoleUsecase(userRoleRepository),
		listUserRolesUsecase:  NewListUserRolesUsecase(userRoleRepository),
		deleteUserRoleUsecase: NewDeleteUserRoleUsecase(userRoleRepository),
		deleteByUserIDUsecase: NewDeleteByUserIDUsecase(userRoleRepository),
		deleteByRoleIDUsecase: NewDeleteByRoleIDUsecase(userRoleRepository),
		countUserRolesUsecase: NewCountUserRolesUsecase(userRoleRepository),
		countByUserIDUsecase:  NewCountByUserIDUsecase(userRoleRepository),
		countByRoleIDUsecase:  NewCountByRoleIDUsecase(userRoleRepository),
		existsUserRoleUsecase: NewExistsUserRoleUsecase(userRoleRepository),
	}
}

func (u *UserRoleUsecaseImpl) Create(ctx context.Context, userRole *entity.UserRole) error {
	return u.createUserRoleUsecase.Execute(ctx, userRole)
}

func (u *UserRoleUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter entity.UserRoleFilter) ([]*entity.UserRole, int64, error) {
	return u.listUserRolesUsecase.Execute(ctx, pagination, filter)
}

func (u *UserRoleUsecaseImpl) Delete(ctx context.Context, userID, roleID string) error {
	return u.deleteUserRoleUsecase.Execute(ctx, userID, roleID)
}

func (u *UserRoleUsecaseImpl) DeleteByUserID(ctx context.Context, userID string) error {
	return u.deleteByUserIDUsecase.Execute(ctx, userID)
}

func (u *UserRoleUsecaseImpl) DeleteByRoleID(ctx context.Context, roleID string) error {
	return u.deleteByRoleIDUsecase.Execute(ctx, roleID)
}

func (u *UserRoleUsecaseImpl) Count(ctx context.Context) (int64, error) {
	return u.countUserRolesUsecase.Execute(ctx)
}

func (u *UserRoleUsecaseImpl) CountByUserID(ctx context.Context, userID string) (int64, error) {
	return u.countByUserIDUsecase.Execute(ctx, userID)
}

func (u *UserRoleUsecaseImpl) CountByRoleID(ctx context.Context, roleID string) (int64, error) {
	return u.countByRoleIDUsecase.Execute(ctx, roleID)
}

func (u *UserRoleUsecaseImpl) Exists(ctx context.Context, userID, roleID string) (bool, error) {
	return u.existsUserRoleUsecase.Execute(ctx, userID, roleID)
}
