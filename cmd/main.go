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
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	permissionServer := permission_service.NewPermissionServer(app.Repos, app.Cacher, app.Helper)
	roleServer := role_service.NewRoleServer(app.Repos)
	rolePermissionServer := role_permission_service.NewRolePermissionServer(app.Repos, app.Helper)
	resourcePermissionServer := resource_permission_service.NewResourcePermissionServer(app.Repos, app.Helper)
	userRoleServer := user_role_service.NewUserRoleServer(app.Repos, app.Helper)
	grpcSrv := grpcservice.NewGRPCServer(env, log, app.Cacher, permissionServer, roleServer, rolePermissionServer, resourcePermissionServer, userRoleServer)
	ctx, cancel := context.WithCancel(context.Background())
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionServer.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to create permissions: " + err.Error())
	}
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
