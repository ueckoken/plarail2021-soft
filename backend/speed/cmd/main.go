package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"ueckoken/plarail2021-soft-speed/internal"
)

const namespace = "plarail2021_speed"

func main() {
	clientConn := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "clients_connections_seconds",
			Help:      "Number of connections handling websocket",
		},
		[]string{},
	)

	controlCommandTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "client_commands_total",
			Help:      "Total client commands",
		},
		[]string{},
	)

	speed := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "speed",
			Help:      "speed",
		},
		[]string{},
	)

	handler := internal.ClientHandler{Clients: &internal.ClientSet{}}
	server := internal.SpeedServer{ClientHandler: &handler, NumberOfClientConnection: clientConn, TotalCLientCommands: controlCommandTotal, Speed: speed}
	server.StartSpeed()
}
