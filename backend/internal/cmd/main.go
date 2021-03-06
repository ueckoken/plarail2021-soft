package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	_ "net/http/pprof"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg/esp32healthcheck"
	"ueckoken/plarail2021-soft-internal/pkg/serveGrpc"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
)

const namespace = "plarailinternal"

func main() {
	env := internal.GetEnv()
	stations, err := station2espIp.NewStations()
	if err != nil {
		log.Fatalln(err)
	}
	esp32HealthCheck := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "esp32_health_seconds",
			Help:      "Esp32 health, if up, then 1",
		},
		[]string{"esp32Addr"},
	)

	go func() {
		fmt.Println("pprof serve at 0.0.0.0:6060")
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	pingHandler := esp32healthcheck.PingHandler{
		Stations:         stations,
		Esp32HealthCheck: esp32HealthCheck,
	}
	grpcServer := serveGrpc.GrpcServer{Stations: stations, Environment: env, PingHandler: pingHandler}
	grpcServer.StartServer()
}
