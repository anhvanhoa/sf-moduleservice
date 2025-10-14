package grpcservice

import (
	"module-service/bootstrap"

	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
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
		},
	)
}
