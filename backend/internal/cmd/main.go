package main

import (
	"log"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg/serveGrpc"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
)

func main() {
	env := internal.GetEnv()
	stations, err := station2espIp.NewStations()
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := serveGrpc.GrpcServer{Stations: stations, Environment: env}
	grpcServer.StartServer()
}
