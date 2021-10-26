package internal

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

const (
	PORT = ":50000"
)

func StartServer() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterControlServer(s, &ControlServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
