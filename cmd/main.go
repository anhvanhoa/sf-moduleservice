package main

import (
	"context"

	"module-service/bootstrap"
	grpcservice "module-service/infrastructure/grpc_service"
	permission_service "module-service/infrastructure/grpc_service/permission"
	resource_permission_service "module-service/infrastructure/grpc_service/resource_permission"
	role_service "module-service/infrastructure/grpc_service/role"
	role_permission_service "module-service/infrastructure/grpc_service/role_permission"
	user_role_service "module-service/infrastructure/grpc_service/user_role"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log
	discoveryConfig := discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGprc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}
	discovery, err := discovery.NewDiscovery(&discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	permissionServer := permission_service.NewPermissionServer(app.Repos, app.Helper)
	roleServer := role_service.NewRoleServer(app.Repos)
	rolePermissionServer := role_permission_service.NewRolePermissionServer(app.Repos, app.Helper)
	resourcePermissionServer := resource_permission_service.NewResourcePermissionServer(app.Repos, app.Helper)
	userRoleServer := user_role_service.NewUserRoleServer(app.Repos, app.Helper)
	grpcSrv := grpcservice.NewGRPCServer(env, log, app.Cacher, permissionServer, roleServer, rolePermissionServer, resourcePermissionServer, userRoleServer)
	ctx, cancel := context.WithCancel(context.Background())
	resources := grpcSrv.GetResources()
	permissions := permissionServer.ConvertResourcesToPermissions(resources)
	if _, err := permissionServer.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to create permissions: " + err.Error())
	}
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
