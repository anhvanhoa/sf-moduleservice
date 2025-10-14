package permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

func (s *permissionService) UpdatePermission(ctx context.Context, req *proto_permission.UpdatePermissionRequest) (*proto_permission.UpdatePermissionResponse, error) {
	permission := s.convertUpdateRequestToEntity(req)
	err := s.permissionUsecase.Update(ctx, permission)
	if err != nil {
		return nil, err
	}
	return &proto_permission.UpdatePermissionResponse{
		Permission: s.convertEntityToProtoPermission(permission),
	}, nil
}

func (s *permissionService) convertUpdateRequestToEntity(req *proto_permission.UpdatePermissionRequest) *entity.Permission {
	return &entity.Permission{
		ID:          req.Id,
		Resource:    req.Resource,
		Action:      req.Action,
		Description: req.Description,
	}
}
