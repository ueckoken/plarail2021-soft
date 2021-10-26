package main

import (
	"ueckoken/plarail2021-soft-external/internal"
)

func main() {
	clientHandler2syncController := make(chan internal.StationState, 16)
	syncController2clientHandler := make(chan internal.StationState, 16)

	envVal := internal.GetEnv()

	httpServer := internal.HttpServer{
		ClientHandler2syncController: clientHandler2syncController,
		SyncController2clientHandler: syncController2clientHandler,
		Environment:                  envVal,
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
