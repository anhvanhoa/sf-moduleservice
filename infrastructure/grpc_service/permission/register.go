package permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) RegisterPermission(ctx context.Context, req *proto_permission.RegisterPermissionRequest) (*proto_permission.RegisterPermissionsResponse, error) {
	permissions := make([]*entity.Permission, len(req.Permissions))
	for i, permission := range req.Permissions {
		permissions[i] = &entity.Permission{
			Resource:    permission.GetResource(),
			Action:      permission.GetAction(),
			Description: permission.GetDescription(),
		}
	}
	err := s.permissionUsecase.Register(ctx, permissions)
	if err != nil {
		return nil, err
	}
	return &proto_permission.RegisterPermissionsResponse{
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
