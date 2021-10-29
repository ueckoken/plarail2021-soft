package internal

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"ueckoken/plarail2021-soft-internal/pkg"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type GrpcServer struct {
	Stations    *pkg.Stations
	Environment *Env
}

func (g *GrpcServer) StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Environment.ExternalSideServer.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	c := ControlServer{
		env:      g.Environment,
		Stations: g.Stations,
	}
	pb.RegisterControlServer(s, &c)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
