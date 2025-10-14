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
	// db := app.DB
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

	permissionServer := permission_service.NewPermissionServer(app.Repos)
	roleServer := role_service.NewRoleServer(app.Repos)
	rolePermissionServer := role_permission_service.NewRolePermissionServer(app.Repos)
	resourcePermissionServer := resource_permission_service.NewResourcePermissionServer(app.Repos)
	userRoleServer := user_role_service.NewUserRoleServer(app.Repos)
	grpcSrv := grpcservice.NewGRPCServer(env, log, permissionServer, roleServer, rolePermissionServer, resourcePermissionServer, userRoleServer)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
