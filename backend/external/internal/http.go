package internal

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	ClientHandler2syncController chan StationState
	SyncController2clientHandler chan StationState
}

type clientsCollection struct {
	Clients []clientChannel
	mtx     sync.Mutex
}

func (h HttpServer) StartServer() {
	clients := clientsCollection{}
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
			fmt.Println("waiting...")
			cChannel := <-clientChannelSend
			fmt.Println("##clients:", len(clients.Clients))
			clients.Clients = append(clients.Clients, cChannel)
			nextClients := []clientChannel{}
			clients.mtx.Lock()
			for _, c := range clients.Clients {
				select {
				case <-c.Done:
					close(c.Done)
					close(c.clientSync)
					continue
				default:
					nextClients = append(nextClients, c)
					//nop
				}
			}
			clients.Clients = nextClients
			clients.mtx.Unlock()
		}
	}()
	for {
		fmt.Println(clients.Clients)
		fmt.Println("goroutine:", runtime.NumGoroutine())
		clients.mtx.Lock()
		for _, c := range clients.Clients {
			select {
			case c.clientSync <- StationState{
				StationID: 0,
				State:     0,
			}:
			default:
				continue
			}
		}
		clients.mtx.Unlock()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("@@@@@@@@@@@@end!!@@@@@@@@@@@@@")
}
