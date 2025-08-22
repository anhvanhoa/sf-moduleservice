package main

import (
	"context"

	"github.com/anhvanhoa/module-service/bootstrap"
	grpcservice "github.com/anhvanhoa/module-service/infrastructure/grpc_service"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log
	db := app.DB
	examService := grpcservice.NewExamService(db, env)
	grpcSrv := grpcservice.NewGRPCServer(env.PORT_GRPC, examService)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
