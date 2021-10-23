package internal

import (
	"fmt"
	"log"
	"net/http"
	pb "ueckoken/plarail2021-soft-external/spec"

	"github.com/gorilla/mux"
)

type HttpServer struct{
	ClientHandler2syncController chan StationKV
	SyncController2clientHandler chan StationKV
}

func (h HttpServer)StartServer() {
	clients := []clientChannel{}
	clientCommand := make(chan pb.RequestSync, 16)
	clientChannelSend := make(chan clientChannel, 16)
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/", handleStatic)
		r.Handle("/ws", clientHandler{ClientCommand: clientCommand, ClientChannelSend: clientChannelSend})
		srv := &http.Server{
			Handler: r,
			Addr:    "127.0.0.1:8000",
			// Good practice: enforce timeouts for servers you create!
		}

		log.Fatal(srv.ListenAndServe())
	}()
	go func() {
		for {
			cChannel := <-clientChannelSend
			clients = append(clients, cChannel)
			nextClients := []clientChannel{}
			for _, c := range clients {
				select {
				case b := <-c.Done:
					if b {
						continue
					}
				default:
					//nop
				}
				nextClients = append(nextClients, c)
			}
			clients = nextClients
		}
	}()
	for {
		fmt.Println(clients)
		for d := range h.SyncController2clientHandler {
			for _, c := range clients {
				c.clientSync <- pb.RequestSync{
					Station: &pb.Stations{StationId: pb.Stations_StationId(d.StationID)},
					State:   pb.RequestSync_State(d.State),
				}
			}
		}
	}
}
