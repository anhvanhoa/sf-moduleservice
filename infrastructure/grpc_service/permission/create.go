package permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) CreatePermission(ctx context.Context, req *proto_permission.CreatePermissionRequest) (*proto_permission.CreatePermissionResponse, error) {
	permission := s.convertRequestCreateToEntity(req)
	err := s.permissionUsecase.Create(ctx, permission)
	if err != nil {
		return nil, err
	}
	return s.convertEntityToResponse(permission), nil
}

func (s *permissionService) convertRequestCreateToEntity(req *proto_permission.CreatePermissionRequest) *entity.Permission {
	p := &entity.Permission{
		Resource:    req.Resource,
		Action:      req.Action,
		Description: req.Description,
	}
	return p
}

func (s *permissionService) convertEntityToResponse(permission *entity.Permission) *proto_permission.CreatePermissionResponse {
	return &proto_permission.CreatePermissionResponse{
		Permission: s.convertEntityToProtoPermission(permission),
	}
}
