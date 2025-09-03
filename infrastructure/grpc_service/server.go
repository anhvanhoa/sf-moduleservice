package grpcservice

import (
	"module-service/bootstrap"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"
	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	grpc_service "github.com/anhvanhoa/service-core/boostrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	moduleService proto_module.ModuleServiceServer,
	moduleChildService proto_module_child.ModuleChildServiceServer,
) *grpc_service.GRPCServer {
	config := &grpc_service.GRPCServerConfig{
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
		IsProduction: env.IsProduction(),
	}
	return grpc_service.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_module.RegisterModuleServiceServer(server, moduleService)
			proto_module_child.RegisterModuleChildServiceServer(server, moduleChildService)
		},
	)
}
