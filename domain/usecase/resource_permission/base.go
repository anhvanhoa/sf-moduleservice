package resource_permission

import (
	"context"
	"module-service/domain/entity"
	"module-service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ResourcePermissionUsecaseI interface {
	Create(ctx context.Context, resourcePermission *entity.ResourcePermission) error
	GetByID(ctx context.Context, id string) (*entity.ResourcePermission, error)
	List(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) (common.PaginationResult[*entity.ResourcePermission], error)
	Update(ctx context.Context, resourcePermission *entity.ResourcePermission) error
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteByResource(ctx context.Context, resourceType, resourceID string) error
	DeleteByUserAndResource(ctx context.Context, userID, resourceType, resourceID string) error
	Count(ctx context.Context) (int64, error)
	CountByUserID(ctx context.Context, userID string) (int64, error)
	CountByResource(ctx context.Context, resourceType, resourceID string) (int64, error)
	Exists(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error)
}

type ResourcePermissionUsecaseImpl struct {
	createResourcePermissionUsecase CreateResourcePermissionUsecase
	getResourcePermissionUsecase    GetResourcePermissionUsecase
	listResourcePermissionsUsecase  ListResourcePermissionsUsecase
	updateResourcePermissionUsecase UpdateResourcePermissionUsecase
	deleteResourcePermissionUsecase DeleteResourcePermissionUsecase
	deleteByUserIDUsecase           DeleteByUserIDUsecase
	deleteByResourceUsecase         DeleteByResourceUsecase
	deleteByUserAndResourceUsecase  DeleteByUserAndResourceUsecase
	countResourcePermissionsUsecase CountResourcePermissionsUsecase
	countByUserIDUsecase            CountByUserIDUsecase
	countByResourceUsecase          CountByResourceUsecase
	existsResourcePermissionUsecase ExistsResourcePermissionUsecase
}

func NewResourcePermissionUsecase(
	resourcePermissionRepository repository.ResourcePermissionRepository,
	helper utils.Helper,
) ResourcePermissionUsecaseI {
	return &ResourcePermissionUsecaseImpl{
		createResourcePermissionUsecase: NewCreateResourcePermissionUsecase(resourcePermissionRepository),
		getResourcePermissionUsecase:    NewGetResourcePermissionUsecase(resourcePermissionRepository),
		listResourcePermissionsUsecase:  NewListResourcePermissionsUsecase(resourcePermissionRepository, helper),
		updateResourcePermissionUsecase: NewUpdateResourcePermissionUsecase(resourcePermissionRepository),
		deleteResourcePermissionUsecase: NewDeleteResourcePermissionUsecase(resourcePermissionRepository),
		deleteByUserIDUsecase:           NewDeleteByUserIDUsecase(resourcePermissionRepository),
		deleteByResourceUsecase:         NewDeleteByResourceUsecase(resourcePermissionRepository),
		deleteByUserAndResourceUsecase:  NewDeleteByUserAndResourceUsecase(resourcePermissionRepository),
		countResourcePermissionsUsecase: NewCountResourcePermissionsUsecase(resourcePermissionRepository),
		countByUserIDUsecase:            NewCountByUserIDUsecase(resourcePermissionRepository),
		countByResourceUsecase:          NewCountByResourceUsecase(resourcePermissionRepository),
		existsResourcePermissionUsecase: NewExistsResourcePermissionUsecase(resourcePermissionRepository),
	}
}

func (u *ResourcePermissionUsecaseImpl) Create(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	return u.createResourcePermissionUsecase.Execute(ctx, resourcePermission)
}

func (u *ResourcePermissionUsecaseImpl) GetByID(ctx context.Context, id string) (*entity.ResourcePermission, error) {
	return u.getResourcePermissionUsecase.Execute(ctx, id)
}

func (u *ResourcePermissionUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter entity.ResourcePermissionFilter) (common.PaginationResult[*entity.ResourcePermission], error) {
	return u.listResourcePermissionsUsecase.Execute(ctx, pagination, filter)
}

func (u *ResourcePermissionUsecaseImpl) Update(ctx context.Context, resourcePermission *entity.ResourcePermission) error {
	return u.updateResourcePermissionUsecase.Execute(ctx, resourcePermission)
}

func (u *ResourcePermissionUsecaseImpl) Delete(ctx context.Context, id string) error {
	return u.deleteResourcePermissionUsecase.Execute(ctx, id)
}

func (u *ResourcePermissionUsecaseImpl) DeleteByUserID(ctx context.Context, userID string) error {
	return u.deleteByUserIDUsecase.Execute(ctx, userID)
}

func (u *ResourcePermissionUsecaseImpl) DeleteByResource(ctx context.Context, resourceType, resourceID string) error {
	return u.deleteByResourceUsecase.Execute(ctx, resourceType, resourceID)
}

func (u *ResourcePermissionUsecaseImpl) DeleteByUserAndResource(ctx context.Context, userID, resourceType, resourceID string) error {
	return u.deleteByUserAndResourceUsecase.Execute(ctx, userID, resourceType, resourceID)
}

func (u *ResourcePermissionUsecaseImpl) Count(ctx context.Context) (int64, error) {
	return u.countResourcePermissionsUsecase.Execute(ctx)
}

func (u *ResourcePermissionUsecaseImpl) CountByUserID(ctx context.Context, userID string) (int64, error) {
	return u.countByUserIDUsecase.Execute(ctx, userID)
}

func (u *ResourcePermissionUsecaseImpl) CountByResource(ctx context.Context, resourceType, resourceID string) (int64, error) {
	return u.countByResourceUsecase.Execute(ctx, resourceType, resourceID)
}

func (u *ResourcePermissionUsecaseImpl) Exists(ctx context.Context, userID, resourceType, resourceID, action string) (bool, error) {
	return u.existsResourcePermissionUsecase.Execute(ctx, userID, resourceType, resourceID, action)
}
