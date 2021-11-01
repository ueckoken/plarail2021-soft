package internal

import (
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"ueckoken/plarail2021-soft-internal/pkg"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type GrpcServer struct {
	Stations    *pkg.Stations
	Environment *Env
}

func (g *GrpcServer) StartServer() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	c := ControlServer{
		env:      g.Environment,
		Stations: g.Stations,
	}
	pb.RegisterControlServer(s, &c)

	// After all your registrations, make sure all the Prometheus metrics are initialized.
	grpcPrometheus.Register(s)
	ServeMetrics(fmt.Sprintf(":%d", g.Environment.ExternalSideServer.MetricsPort))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Environment.ExternalSideServer.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func ServeMetrics(promAddr string) {
	mux := http.NewServeMux()
	// Enable histogram
	grpcPrometheus.EnableHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Prometheus metrics bind address", promAddr)
		log.Fatal(http.ListenAndServe(promAddr, mux))
	}()
}
