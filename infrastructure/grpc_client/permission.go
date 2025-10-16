package gc

import (
	"context"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/grpc_client"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
)

type PermissionClient struct {
	client   *grpc_client.Client
	pservice proto_permission.PermissionServiceClient
}

func NewPermissionClient(client *grpc_client.Client) *PermissionClient {
	return &PermissionClient{
		client:   client,
		pservice: proto_permission.NewPermissionServiceClient(client.GetConnection()),
	}
}

func (c *PermissionClient) CreateManyPermission(
	ctx context.Context,
	resources grpc_server.InfoResources,
) error {
	permissions := c.convertResourcesToPermissions(resources)
	_, err := c.pservice.CreatePermissions(ctx, &proto_permission.CreatePermissionsRequest{
		Permissions: permissions,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *PermissionClient) convertResourcesToPermissions(resources grpc_server.InfoResources) []*proto_permission.CreatePermissionRequest {
	permissions := make([]*proto_permission.CreatePermissionRequest, 0)
	for resource, methods := range resources {
		for _, method := range methods {
			permissions = append(permissions, &proto_permission.CreatePermissionRequest{
				Resource: resource,
				Action:   method,
			})
		}
	}
	return permissions
}
