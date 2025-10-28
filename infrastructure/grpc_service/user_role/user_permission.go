package user_role_service

import (
	"context"

	"github.com/anhvanhoa/service-core/domain/user_context"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) GetUserPermissions(ctx context.Context, req *proto_user_role.GetUserPermissionsRequest) (*proto_user_role.GetUserPermissionsResponse, error) {
	userPermission, err := s.userRoleUsecase.GetUserPermissions(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return s.convertEntityToProtoUserPermissions(userPermission), nil
}

func (s *userRoleService) convertEntityToProtoUserPermissions(userPermission user_context.UserContext) *proto_user_role.GetUserPermissionsResponse {
	protoPermissions := make([]*proto_user_role.Permission, len(userPermission.Permissions))
	for i, permission := range userPermission.Permissions {
		protoPermissions[i] = &proto_user_role.Permission{
			Resource: permission.Resource,
			Action:   permission.Action,
		}
	}
	protoScopes := make([]*proto_user_role.Scope, len(userPermission.Scopes))
	for i, scope := range userPermission.Scopes {
		protoScopes[i] = &proto_user_role.Scope{
			Resource:     scope.Resource,
			ResourceData: scope.ResourceData,
			Action:       scope.Action,
		}
	}
	return &proto_user_role.GetUserPermissionsResponse{
		UserId:      userPermission.UserID,
		Roles:       userPermission.Roles,
		Permissions: protoPermissions,
		Scopes:      protoScopes,
	}
}
