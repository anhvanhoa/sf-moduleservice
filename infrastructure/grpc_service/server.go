package grpcservice

import (
	"context"
	"fmt"
	"net"

	"module-service/bootstrap"
	"module-service/domain/service/logger"
	proto "module-service/proto/gen/module/v1"

	"buf.build/go/protovalidate"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	server *grpc.Server
	port   string
	log    logger.Log
}

func NewGRPCServer(
	env *bootstrap.Env,
	log logger.Log,
	moduleService proto.ModuleServiceServer,
	moduleChildService proto.ModuleChildServiceServer,
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

	proto.RegisterModuleServiceServer(server, moduleService)
	proto.RegisterModuleChildServiceServer(server, moduleChildService)

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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.log.Error("failed to listen", err)
	}

	s.log.Info(fmt.Sprintf("gRPC server starting on port %s", s.port))

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
