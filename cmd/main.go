package main

import (
	"context"

	"module-service/bootstrap"
	grpcservice "module-service/infrastructure/grpc_service"
	moduleservice "module-service/infrastructure/grpc_service/module"
	modulechildservice "module-service/infrastructure/grpc_service/module_child"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log
	db := app.DB
	moduleService := moduleservice.NewModuleService(db, env)
	moduleChildService := modulechildservice.NewModuleChildService(db, env)
	grpcSrv := grpcservice.NewGRPCServer(env.PORT_GRPC, log, moduleService, moduleChildService)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
