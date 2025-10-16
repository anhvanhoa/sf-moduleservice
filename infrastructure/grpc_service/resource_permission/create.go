package resource_permission_service

import (
	"context"
	"module-service/domain/entity"

	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
)

func (s *resourcePermissionService) CreateResourcePermission(ctx context.Context, req *proto_resource_permission.CreateResourcePermissionRequest) (*proto_resource_permission.CreateResourcePermissionResponse, error) {
	resourcePermission := s.convertRequestCreateToEntity(req)
	err := s.resourcePermissionUsecase.Create(ctx, resourcePermission)
	if err != nil {
		return nil, err
	}
	return s.convertEntityToResponse(resourcePermission), nil
}

func (s *resourcePermissionService) convertRequestCreateToEntity(req *proto_resource_permission.CreateResourcePermissionRequest) *entity.ResourcePermission {
	return &entity.ResourcePermission{
		UserID:       req.UserId,
		ResourceType: req.ResourceType,
		ResourceData: req.ResourceData,
		Action:       req.Action,
	}
}

func (s *resourcePermissionService) convertEntityToResponse(rp *entity.ResourcePermission) *proto_resource_permission.CreateResourcePermissionResponse {
	return &proto_resource_permission.CreateResourcePermissionResponse{
		ResourcePermission: s.convertEntityToProtoResourcePermission(rp),
	}
}
