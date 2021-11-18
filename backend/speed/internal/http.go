package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-speed/pkg/storeSpeed"
)

type SpeedServer struct {
	ClientHandler *ClientHandler
}

func (s SpeedServer) StartSpeed() {
	reg := make(chan Client)
	change := make(chan storeSpeed.TrainConf)
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	s.ClientHandler.upgrader = upgrader
	s.ClientHandler.ClientNotification = reg
	s.ClientHandler.ClientCommand = change
	go s.RegisterClient(reg)
	go s.HandleChange(change)
	go s.UnregisterClient()
	http.Handle("/speed", s.ClientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ClientHandler struct {
	Clients            *ClientSet
	upgrader           websocket.Upgrader
	ClientNotification chan Client
	ClientCommand      chan storeSpeed.TrainConf
}

type ClientSet struct {
	mtx     sync.Mutex
	clients []Client
}
type Client struct {
	notifier ClientNotifier
}

type ClientNotifier struct {
	Notifier   chan storeSpeed.TrainConf
	Unregister chan struct{}
}

func (s *SpeedServer) RegisterClient(cn chan Client) {
	for {
		select {
		case n := <-cn:
			s.ClientHandler.Clients.mtx.Lock()
			s.ClientHandler.Clients.clients = append(s.ClientHandler.Clients.clients, n)
			s.ClientHandler.Clients.mtx.Unlock()
		}
	}
}

func (s *SpeedServer) HandleChange(cn chan storeSpeed.TrainConf) {
	for {
		select {
		case c := <-cn:
			//this should be sorted from old to new
			for _, client := range s.ClientHandler.Clients.clients {
				select {
				case client.notifier.Notifier <- c:
				default:
					fmt.Println("buffer is full...")
				}
			}
		}
	}
}

func (s *SpeedServer) UnregisterClient() {
	for {
		s.ClientHandler.Clients.mtx.Lock()
		var deletionList []int
		for i, c := range s.ClientHandler.Clients.clients {
			select {
			case <-c.notifier.Unregister:
				deletionList = append(deletionList, i)
			default:
				continue
			}
		}
		s.ClientHandler.Clients.deleteClient(deletionList)
		s.ClientHandler.Clients.mtx.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (cl *ClientSet) deleteClient(deletion []int) {
	var tmp []Client
	for i, c := range cl.clients {
		if !contain(deletion, i) {
			tmp = append(tmp, c)
		}
	}
	cl.clients = tmp
}

func contain(list []int, data int) bool {
	for _, l := range list {
		if l == data {
			return true
		}
	}
	return false
}
