package resource_permission_service

import (
	"module-service/domain/entity"
	"module-service/domain/usecase/resource_permission"
	"module-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/utils"
	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

type resourcePermissionService struct {
	proto_resource_permission.UnimplementedResourcePermissionServiceServer
	resourcePermissionUsecase resource_permission.ResourcePermissionUsecaseI
}

func NewResourcePermissionServer(repos repo.Repositories, helper utils.Helper) proto_resource_permission.ResourcePermissionServiceServer {
	resourcePermissionRepo := repos.ResourcePermissionRepository()
	resourcePermissionUC := resource_permission.NewResourcePermissionUsecase(resourcePermissionRepo, helper)
	return &resourcePermissionService{
		resourcePermissionUsecase: resourcePermissionUC,
	}
}

func (s *resourcePermissionService) convertEntityToProtoResourcePermission(rp *entity.ResourcePermission) *proto_resource_permission.ResourcePermission {
	return &proto_resource_permission.ResourcePermission{
		Id:           rp.ID,
		UserId:       rp.UserID,
		ResourceType: rp.ResourceType,
		ResourceData: rp.ResourceData,
		Action:       rp.Action,
	}
}
