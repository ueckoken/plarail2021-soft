package main

import (
	"ueckoken/plarail2021-soft-speed/internal"
)

func main() {
	handler := internal.ClientHandler{Clients: &internal.ClientSet{}}
	server := internal.SpeedServer{ClientHandler: &handler}
	server.StartSpeed()
}
