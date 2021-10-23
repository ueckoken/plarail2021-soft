package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpServer struct{
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
}

func (h HttpServer)StartServer() {
	clients := []clientChannel{}
	clientCommand := make(chan StationState, 16)
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
				c.clientSync <- StationState{
					StationID: d.StationID,
					State:   d.State,
				}
			}
		}
	}
}
