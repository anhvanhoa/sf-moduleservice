package permission_service

import (
	"module-service/domain/entity"
	"module-service/domain/usecase/permission"
	"module-service/infrastructure/repo"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/utils"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

type PermissionService interface {
	ConvertResourcesToPermissions(resources grpc_server.InfoResources) *proto_permission.RegisterPermissionRequest
	proto_permission.PermissionServiceServer
}

type permissionService struct {
	proto_permission.UnimplementedPermissionServiceServer
	permissionUsecase permission.PermissionUsecaseI
}

func NewPermissionServer(repos repo.Repositories, helper utils.Helper) PermissionService {
	permissionRepo := repos.PermissionRepository()
	permissionUC := permission.NewPermissionUsecase(permissionRepo, helper)
	return &permissionService{
		permissionUsecase: permissionUC,
	}
}
func (s *permissionService) convertEntityToProtoPermission(permission *entity.Permission) *proto_permission.Permission {
	return &proto_permission.Permission{
		Id:          permission.ID,
		Resource:    permission.Resource,
		Action:      permission.Action,
		Description: permission.Description,
	}
}

func (s *permissionService) ConvertResourcesToPermissions(resources grpc_server.InfoResources) *proto_permission.RegisterPermissionRequest {
	permissions := make([]*proto_permission.CreatePermissionRequest, 0)
	for resource, methods := range resources {
		for _, method := range methods {
			permissions = append(permissions, &proto_permission.CreatePermissionRequest{
				Resource: resource,
				Action:   method,
			})
		}
	}
	return &proto_permission.RegisterPermissionRequest{
		Permissions: permissions,
	}
}
