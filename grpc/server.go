package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	listener net.Listener
	server   *grpc.Server
	health   *health.Server
}

type GRPCServerRegister interface {
	Register(*grpc.Server)
}

func (srv *GRPCServer) NewServer(address string) *GRPCServer {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)

	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(s)

	srv.listener = lis
	srv.server = s
	srv.health = healthServer

	return srv
}

func (srv *GRPCServer) Register(interfaces ...GRPCServerRegister) *GRPCServer {
	for _, itf := range interfaces {
		itf.Register(srv.server)
	}

	return srv
}

func (srv *GRPCServer) Serve() error {
	return srv.server.Serve(srv.listener)
}

func (srv *GRPCServer) SetHealth(isHealthy bool) {
	if srv.health == nil {
		return
	}

	if isHealthy {
		srv.health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	} else {
		srv.health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	}
}
