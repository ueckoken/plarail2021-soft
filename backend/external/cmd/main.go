package main

import (
	"ueckoken/plarail2021-soft-external/internal"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	clientHandler2syncController := make(chan internal.StationState)
	syncController2clientHandler := make(chan internal.StationState)

	envVal := internal.GetEnv()

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
	syncController := internal.SyncController{
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
	}
	go httpServer.StartServer()
	go syncController.StartSyncController()
	for {
	}
}
