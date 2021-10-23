package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/clientSync"

	"github.com/gorilla/mux"
)

func StartServer() {
	clients := []clientChannel{}
	clientCommand := make(chan clientSync.SingleState, 16)
	clientChannelSend := make(chan clientChannel, 16)
	go func() {
		r := mux.NewRouter()
    r.HandleFunc("/", handleStatic)
		r.Handle("/ws", clientHandler{ClientCommand: clientCommand, ClientChannelSend: clientChannelSend})
		srv := &http.Server{
			Handler: r,
			Addr:    "127.0.0.1:8000",
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
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
    for _,c := range clients{
      c.clientSync <- clientSync.ClientSync{State: []clientSync.SingleState{clientSync.SingleState{Name: "id", OnOff: false}}}
    }
		time.Sleep(1 * time.Second)
	}
}
