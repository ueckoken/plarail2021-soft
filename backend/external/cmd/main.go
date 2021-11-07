package main

import (
	"ueckoken/plarail2021-soft-external/internal"
	"ueckoken/plarail2021-soft-external/pkg/envStore"
	"ueckoken/plarail2021-soft-external/pkg/syncController"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	v := syncController.NewRouteValidator()
	println("%v", v)
	clientHandler2syncController := make(chan syncController.StationState)
	syncController2clientHandler := make(chan syncController.StationState)

	envVal := envStore.GetEnv()

	clientconn := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "plarailexternal_client_connections",
			Help: "Number of connections handling websocket",
		},
		[]string{"client"},
	)

	httpServer := internal.HttpServer{
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
		NumberOfClientConnection:     clientconn,
	}
	syncController := syncController.SyncController{
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
	}
	go httpServer.StartServer()
	syncController.StartSyncController()
}
