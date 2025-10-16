package resource_permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) UpdateResourcePermission(ctx context.Context, req *proto_resource_permission.UpdateResourcePermissionRequest) (*proto_resource_permission.UpdateResourcePermissionResponse, error) {
	resourcePermission := s.convertUpdateRequestToEntity(req)
	err := s.resourcePermissionUsecase.Update(ctx, resourcePermission)
	if err != nil {
		return nil, err
	}
	return &proto_resource_permission.UpdateResourcePermissionResponse{
		ResourcePermission: s.convertEntityToProtoResourcePermission(resourcePermission),
	}, nil
}

func (s *resourcePermissionService) convertUpdateRequestToEntity(req *proto_resource_permission.UpdateResourcePermissionRequest) *entity.ResourcePermission {
	return &entity.ResourcePermission{
		ID:           req.Id,
		UserID:       req.UserId,
		ResourceType: req.ResourceType,
		ResourceData: req.ResourceData,
		Action:       req.Action,
	}
}
