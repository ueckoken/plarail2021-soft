package internal

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "ueckoken/plarail2021-soft-builtin/spec"
)

type GrpcServer struct {
	Port int
}

func NewGrpcServer(servPort int) *GrpcServer {
	return &GrpcServer{Port: servPort}
}
func (g *GrpcServer) StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	c := NewServeSpeedControl()
	pb.RegisterSpeedServer(s, c)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
