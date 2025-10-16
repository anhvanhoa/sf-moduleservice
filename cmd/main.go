package main

import (
	"context"

	"module-service/bootstrap"
	gc "module-service/infrastructure/grpc_client"
	grpcservice "module-service/infrastructure/grpc_service"
	permission_service "module-service/infrastructure/grpc_service/permission"
	resource_permission_service "module-service/infrastructure/grpc_service/resource_permission"
	role_service "module-service/infrastructure/grpc_service/role"
	role_permission_service "module-service/infrastructure/grpc_service/role_permission"
	user_role_service "module-service/infrastructure/grpc_service/user_role"

	"github.com/anhvanhoa/service-core/domain/discovery"
	"github.com/anhvanhoa/service-core/domain/grpc_client"
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

	clientFactory := grpc_client.NewClientFactory(env.GrpcClients...)
	permissionClient := gc.NewPermissionClient(clientFactory.GetClient(env.AddressPermission))

	permissionServer := permission_service.NewPermissionServer(app.Repos, app.Helper)
	roleServer := role_service.NewRoleServer(app.Repos)
	rolePermissionServer := role_permission_service.NewRolePermissionServer(app.Repos, app.Helper)
	resourcePermissionServer := resource_permission_service.NewResourcePermissionServer(app.Repos, app.Helper)
	userRoleServer := user_role_service.NewUserRoleServer(app.Repos, app.Helper)
	grpcSrv := grpcservice.NewGRPCServer(env, log, app.Cacher, permissionServer, roleServer, rolePermissionServer, resourcePermissionServer, userRoleServer)
	ctx, cancel := context.WithCancel(context.Background())
	resources := grpcSrv.GetResources()
	if err := permissionClient.CreateManyPermission(ctx, resources); err != nil {
		log.Fatal("Failed to create many permission: " + err.Error())
	}
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
