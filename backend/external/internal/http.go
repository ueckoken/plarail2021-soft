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
	Environment                  *Env
}

type clientsCollection struct {
	Clients []clientChannel
	mtx     sync.Mutex
}

func (h HttpServer) StartServer() {
	clients := clientsCollection{}
	clientChannelSend := make(chan clientChannel)
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/", handleStatic)
		r.Handle("/ws", clientHandler{ClientCommand: h.ClientHandler2syncController, ClientChannelSend: clientChannelSend})
		srv := &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf("0.0.0.0:%d", h.Environment.ClientSideServer.Port),
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
	go func() {
		for {
			h.SyncController2clientHandler <- StationState{}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Println(clients.Clients)
		fmt.Println("goroutine:", runtime.NumGoroutine())
		for d := range h.SyncController2clientHandler {
			fmt.Println(clients.Clients)
			clients.mtx.Lock()
			for _, c := range clients.Clients {
				select {
				case c.clientSync <- d:
				default:
					continue
				}
			}
			clients.mtx.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}
