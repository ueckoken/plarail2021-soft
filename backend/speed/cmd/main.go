package main

import (
	"ueckoken/plarail2021-soft-speed/internal"
	"ueckoken/plarail2021-soft-speed/pkg/train2IP"
)

func main() {
	trainInfo := train2IP.GetTable()
	handler := internal.ClientHandler{Clients: &internal.ClientSet{}, TrainName2Id: trainInfo}
	server := internal.SpeedServer{ClientHandler: &handler}
	server.StartSpeed()
}
