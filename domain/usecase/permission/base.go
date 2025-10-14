package permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type PermissionUsecaseI interface {
	Create(ctx context.Context, permission *entity.Permission) error
	GetByID(ctx context.Context, id string) (*entity.Permission, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) (common.PaginationResult[*entity.Permission], error)
	Update(ctx context.Context, permission *entity.Permission) error
	Delete(ctx context.Context, id string) error
	DeleteByResourceAndAction(ctx context.Context, resource, action string) error
	CountByResource(ctx context.Context, resource string) (int64, error)
}

type PermissionUsecaseImpl struct {
	createPermissionUsecase          CreatePermissionUsecase
	getPermissionUsecase             GetPermissionUsecase
	listPermissionsUsecase           ListPermissionsUsecase
	updatePermissionUsecase          UpdatePermissionUsecase
	deletePermissionUsecase          DeletePermissionUsecase
	deleteByResourceAndActionUsecase DeleteByResourceAndActionUsecase
	countByResourceUsecase           CountByResourceUsecase
}

func NewPermissionUsecase(permissionRepository repository.PermissionRepository, helper utils.Helper) PermissionUsecaseI {
	return &PermissionUsecaseImpl{
		createPermissionUsecase:          NewCreatePermissionUsecase(permissionRepository),
		getPermissionUsecase:             NewGetPermissionUsecase(permissionRepository),
		listPermissionsUsecase:           NewListPermissionsUsecase(permissionRepository, helper),
		updatePermissionUsecase:          NewUpdatePermissionUsecase(permissionRepository),
		deletePermissionUsecase:          NewDeletePermissionUsecase(permissionRepository),
		deleteByResourceAndActionUsecase: NewDeleteByResourceAndActionUsecase(permissionRepository),
		countByResourceUsecase:           NewCountByResourceUsecase(permissionRepository),
	}
}

func (u *PermissionUsecaseImpl) Create(ctx context.Context, permission *entity.Permission) error {
	return u.createPermissionUsecase.Execute(ctx, permission)
}

func (u *PermissionUsecaseImpl) GetByID(ctx context.Context, id string) (*entity.Permission, error) {
	return u.getPermissionUsecase.Execute(ctx, id)
}

func (u *PermissionUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter entity.PermissionFilter) (common.PaginationResult[*entity.Permission], error) {
	return u.listPermissionsUsecase.Execute(ctx, pagination, filter)
}

func (u *PermissionUsecaseImpl) Update(ctx context.Context, permission *entity.Permission) error {
	return u.updatePermissionUsecase.Execute(ctx, permission)
}

func (u *PermissionUsecaseImpl) Delete(ctx context.Context, id string) error {
	return u.deletePermissionUsecase.Execute(ctx, id)
}

func (u *PermissionUsecaseImpl) DeleteByResourceAndAction(ctx context.Context, resource, action string) error {
	return u.deleteByResourceAndActionUsecase.Execute(ctx, resource, action)
}

func (u *PermissionUsecaseImpl) CountByResource(ctx context.Context, resource string) (int64, error) {
	return u.countByResourceUsecase.Execute(ctx, resource)
}
