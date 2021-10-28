package main

import (
	"log"
	"ueckoken/plarail2021-soft-internal/internal"
	"ueckoken/plarail2021-soft-internal/pkg"
)

func main() {
	env := internal.GetEnv()
	stations, err := pkg.NewStations()
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := internal.GrpcServer{Stations: stations, Environment: env}
	grpcServer.StartServer()
}
