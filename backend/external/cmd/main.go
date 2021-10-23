package main

import (
	"ueckoken/plarail2021-soft-external/internal"
)

func main() {
	clientHandler2syncController := make(chan internal.StationKV, 16)
	syncController2clientHandler := make(chan internal.StationKV, 16)
	httpServer := internal.HttpServer{ClientHandler2syncController: clientHandler2syncController, SyncController2clientHandler: syncController2clientHandler}
	syncController := internal.SyncController{ClientHandler2syncController: clientHandler2syncController, SyncController2clientHandler: syncController2clientHandler}
	go httpServer.StartServer()
	go syncController.StartSyncController()
	for {
	}
}
