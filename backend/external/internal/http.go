package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"
	pb "ueckoken/plarail2021-soft-external/spec"

	"github.com/gorilla/mux"
)

func StartServer() {
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
		for _, c := range clients {
			c.clientSync <- pb.RequestSync{Name: "id", State: pb.RequestSync_ON}
		}
		time.Sleep(1 * time.Second)
	}
}
