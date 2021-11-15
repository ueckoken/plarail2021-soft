package serveGrpc

import (
	"fmt"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg/esp32healthcheck"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
	pb "ueckoken/plarail2021-soft-internal/spec"
)

type GrpcServer struct {
	Stations    station2espIp.Stations
	Environment *internal.Env
	PingHandler esp32healthcheck.PingHandler
}

func (g *GrpcServer) StartServer() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcPrometheus.UnaryServerInterceptor),
	)
	c := ControlServer{
		Stations: g.Stations,
		client:   &http.Client{Timeout: g.Environment.NodeConnection.Timeout},
	}
	pb.RegisterControlServer(s, &c)

	// After all your registrations, make sure all the Prometheus metrics are initialized.
	grpcPrometheus.Register(s)
	prometheus.MustRegister(g.PingHandler.Esp32HealthCheck)
	go g.PingHandler.Start()
	go ServeMetrics(fmt.Sprintf(":%d", g.Environment.ExternalSideServer.MetricsPort))

	port := fmt.Sprintf(":%d", g.Environment.ExternalSideServer.Port)
	log.Println("gRPC serve at", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serveGrpc: %v", err)
	}
}

func ServeMetrics(promAddr string) {
	mux := http.NewServeMux()
	// Enable histogram
	grpcPrometheus.EnableHandlingTimeHistogram()
	mux.Handle("/metrics", promhttp.Handler())
	fmt.Println("Prometheus metrics bind address", promAddr)
	log.Fatal(http.ListenAndServe(promAddr, mux))
}
