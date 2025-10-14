package grpcservice

import (
	"module-service/bootstrap"

	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
	proto_role "github.com/anhvanhoa/sf-proto/gen/role/v1"
	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"

	"github.com/anhvanhoa/service-core/domain/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	permissionServer proto_permission.PermissionServiceServer,
	roleServer proto_role.RoleServiceServer,
	rolePermissionServer proto_role_permission.RolePermissionServiceServer,
	resourcePermissionServer proto_resource_permission.ResourcePermissionServiceServer,
	userRoleServer proto_user_role.UserRoleServiceServer,
) *grpc_service.GRPCServer {
	config := &grpc_service.GRPCServerConfig{
		PortGRPC:     env.PortGprc,
		NameService:  env.NameService,
		IsProduction: env.IsProduction(),
	}
	return grpc_service.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_permission.RegisterPermissionServiceServer(server, permissionServer)
			proto_role.RegisterRoleServiceServer(server, roleServer)
			proto_role_permission.RegisterRolePermissionServiceServer(server, rolePermissionServer)
			proto_resource_permission.RegisterResourcePermissionServiceServer(server, resourcePermissionServer)
			proto_user_role.RegisterUserRoleServiceServer(server, userRoleServer)
		},
	)
}
