package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/clientHandler"
	"ueckoken/plarail2021-soft-external/pkg/envStore"
	"ueckoken/plarail2021-soft-external/pkg/syncController"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HttpServer struct {
	ClientHandler2syncController chan syncController.StationState
	SyncController2clientHandler chan syncController.StationState
	Environment                  *envStore.Env
	NumberOfClientConnection     *prometheus.GaugeVec
	TotalClientConnection        *prometheus.CounterVec
	TotalCLientCommands          *prometheus.CounterVec
}

type clientsCollection struct {
	Clients []clientHandler.ClientChannel
	mtx     sync.Mutex
}

func (h HttpServer) StartServer() {
	clients := clientsCollection{}
	clientChannelSend := make(chan clientHandler.ClientChannel)
	go func() {
		r := mux.NewRouter()
		prometheus.MustRegister(h.NumberOfClientConnection)
		prometheus.MustRegister(h.TotalClientConnection)
		prometheus.MustRegister(h.TotalCLientCommands)
		r.HandleFunc("/", clientHandler.HandleStatic)
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		r.Handle("/ws", clientHandler.ClientHandler{Upgrader: upgrader, ClientCommand: h.ClientHandler2syncController, ClientChannelSend: clientChannelSend})
		r.Handle("/metrics", promhttp.Handler())
		srv := &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf("0.0.0.0:%d", h.Environment.ClientSideServer.Port),
			// Good practice: enforce timeouts for servers you create!
		}

		log.Fatal(srv.ListenAndServe())
	}()
	go func() {
		for {
			cChannel := <-clientChannelSend
			clients.Clients = append(clients.Clients, cChannel)
			h.TotalClientConnection.With(prometheus.Labels{}).Inc()
			nextClients := []clientHandler.ClientChannel{}
			clients.mtx.Lock()
			for _, c := range clients.Clients {
				select {
				case <-c.Done:
					close(c.Done)
					close(c.ClientSync)
					continue
				default:
					nextClients = append(nextClients, c)
					//nop
				}
			}
			clients.Clients = nextClients
			h.NumberOfClientConnection.With(prometheus.Labels{}).Set(float64(len(clients.Clients)))
			clients.mtx.Unlock()
		}
	}()
	go func() {
		for {
			h.SyncController2clientHandler <- syncController.StationState{}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		for d := range h.SyncController2clientHandler {
			h.TotalCLientCommands.With(prometheus.Labels{}).Inc()
			clients.mtx.Lock()
			for _, c := range clients.Clients {
				select {
				case c.ClientSync <- d:
				default:
					continue
				}
			}
			clients.mtx.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}
