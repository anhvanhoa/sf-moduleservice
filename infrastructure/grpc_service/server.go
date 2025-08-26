package grpcservice

import (
	"context"
	"fmt"
	"net"

	"module-service/bootstrap"
	"module-service/domain/service/logger"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"
	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"

	"buf.build/go/protovalidate"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	port   int
	log    logger.Log
}

func NewGRPCServer(
	env *bootstrap.Env,
	log logger.Log,
	moduleService proto_module.ModuleServiceServer,
	moduleChildService proto_module_child.ModuleChildServiceServer,
) *GRPCServer {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			protovalidate_middleware.UnaryServerInterceptor(validator),
			LoggingInterceptor(log),
		),
	)

	proto_module.RegisterModuleServiceServer(server, moduleService)
	proto_module_child.RegisterModuleChildServiceServer(server, moduleChildService)

	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthSrv)

	healthSrv.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	healthSrv.SetServingStatus(env.NAME_SERVICE, grpc_health_v1.HealthCheckResponse_SERVING)

	if !env.IsProduction() {
		reflection.Register(server)
		log.Info("Reflection enabled")
	}

	return &GRPCServer{
		server: server,
		port:   env.PORT_GRPC,
		log:    log,
	}
}

func (s *GRPCServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		s.log.Error("failed to listen", err)
	}

	s.log.Info(fmt.Sprintf("gRPC server starting on port %d", s.port))

	go func() {
		if err := s.server.Serve(lis); err != nil {
			s.log.Error("failed to serve", err)
		}
	}()

	<-ctx.Done()

	s.log.Info("Shutting down gRPC server...")
	s.server.GracefulStop()

	return nil
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}

func (s *GRPCServer) GetServer() *grpc.Server {
	return s.server
}
