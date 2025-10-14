package main

import (
	"context"

	"module-service/bootstrap"
	grpcservice "module-service/infrastructure/grpc_service"

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

	grpcSrv := grpcservice.NewGRPCServer(env, log)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
