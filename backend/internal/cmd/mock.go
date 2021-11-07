package main

import (
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"log"
	"net"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg/grpcMock"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	e := internal.GetEnv()
	m := grpcMock.GrpcMock{}
	pb.RegisterControlServer(s, &m)
	grpcPrometheus.Register(s)
	internal.ServeMetrics(fmt.Sprintf(":%d", e.ExternalSideServer.MetricsPort))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", e.ExternalSideServer.Port))
	if err != nil {
		log.Fatalf("listen failed")
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
