package permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) CreateManyPermission(ctx context.Context, req *proto_permission.CreatePermissionsRequest) (*proto_permission.CreatePermissionsResponse, error) {
	permissions := make([]*entity.Permission, len(req.Permissions))
	for i, permission := range req.Permissions {
		permissions[i] = &entity.Permission{
			Resource:    permission.Resource,
			Action:      permission.Action,
			Description: permission.Description,
		}
	}
	err := s.permissionUsecase.CreateMany(ctx, permissions)
	if err != nil {
		return nil, err
	}
	return &proto_permission.CreatePermissionsResponse{
		Permissions: s.convertEntityToProtoPermissions(permissions),
	}, nil
}

func (s *permissionService) convertEntityToProtoPermissions(permissions []*entity.Permission) []*proto_permission.Permission {
	protoPermissions := make([]*proto_permission.Permission, len(permissions))
	for i, permission := range permissions {
		protoPermissions[i] = s.convertEntityToProtoPermission(permission)
	}
	return protoPermissions
}
