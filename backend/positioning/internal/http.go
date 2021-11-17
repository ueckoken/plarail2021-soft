package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"ueckoken/plarail2021-soft-positioning/pkg/gormHandler"
	"ueckoken/plarail2021-soft-positioning/pkg/positionReceiver"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"
)

type PositionReceiver struct {
	db      gormHandler.SQLHandler
	status  ApplicationStatus
	clients *ClientSet
}

func NewPositionReceiver(db gormHandler.SQLHandler, status ApplicationStatus) PositionReceiver {
	cls := ClientSet{}
	return PositionReceiver{db: db, status: status, clients: &cls}
}

func (pos PositionReceiver) StartPositionReceiver() {
	c := make(chan trainState.State)
	p := positionReceiver.NewPositionReceiverHandler(c)
	n := make(chan ClientNotifier)
	h := ClientHandler{ClientNotification: n}
	go pos.RegisterClient(n)
	go pos.HandleChange(c)
	go pos.UnregisterClient()
	http.Handle("/registerPosition", p)
	http.Handle("/subscribePosition", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type ClientSet struct {
	mtx     sync.Mutex
	clients []Client
}

type Client struct {
	notifier ClientNotifier
}

func (pos *PositionReceiver) RegisterClient(cn chan ClientNotifier) {
	for {
		select {
		case n := <-cn:
			pos.clients.mtx.Lock()
			pos.clients.clients = append(pos.clients.clients, Client{n})
			pos.clients.mtx.Unlock()
		}
	}
}

func (pos *PositionReceiver) HandleChange(cn chan trainState.State) {
	for {
		select {
		case c := <-cn:
			for _, client := range pos.clients.clients {
				select {
				case client.notifier.Notifier <- c:
				default:
					fmt.Println("buffer is full...")
				}
			}
		}
	}
}

func (pos *PositionReceiver) UnregisterClient() {
	for {
		pos.clients.mtx.Lock()
		var deletionList []int
		for i, c := range pos.clients.clients {
			select {
			case <-c.notifier.Unregister:
				deletionList = append(deletionList, i)
			default:
				continue
			}
		}
		pos.clients.deleteClient(deletionList)
		pos.clients.mtx.Unlock()
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
