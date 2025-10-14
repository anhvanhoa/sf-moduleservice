package role_permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
)

func (s *rolePermissionService) CreateRolePermission(ctx context.Context, req *proto_role_permission.CreateRolePermissionRequest) (*proto_role_permission.CreateRolePermissionResponse, error) {
	rolePermission := s.convertRequestCreateToEntity(req)
	err := s.rolePermissionUsecase.Create(ctx, rolePermission)
	if err != nil {
		return nil, err
	}
	return s.convertEntityToResponse(rolePermission), nil
}

func (s *rolePermissionService) convertRequestCreateToEntity(req *proto_role_permission.CreateRolePermissionRequest) *entity.RolePermission {
	return &entity.RolePermission{
		RoleID:       req.RoleId,
		PermissionID: req.PermissionId,
	}
}

func (s *rolePermissionService) convertEntityToResponse(rp *entity.RolePermission) *proto_role_permission.CreateRolePermissionResponse {
	return &proto_role_permission.CreateRolePermissionResponse{
		RolePermission: s.convertEntityToProtoRolePermission(rp),
	}
}
