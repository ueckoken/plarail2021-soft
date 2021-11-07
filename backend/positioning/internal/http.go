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
	db gormHandler.SQLHandler
}

func NewPositionReceiver(db gormHandler.SQLHandler) PositionReceiver {
	return PositionReceiver{db}
}

func (pos PositionReceiver) StartPositionReceiver() {
	c := make(chan trainState.State)
	p := positionReceiver.NewPositionReceiverHandler(c, pos.db)
	http.Handle("/registerPosition", p)
	n := make(chan ClientNotifier)
	h := ClientHandler{ClientNotification: n}
	cls := ClientSet{}
	go cls.RegisterClient(n)
	go cls.HandleChange(c)
	go cls.UnregisterClient()
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

func (cl *ClientSet) RegisterClient(cn chan ClientNotifier) {
	for {
		select {
		case n := <-cn:
			cl.mtx.Lock()
			cl.clients = append(cl.clients, Client{n})
			cl.mtx.Unlock()
		}
	}
}

func (cl *ClientSet) HandleChange(cn chan trainState.State) {
	for {
		select {
		case c := <-cn:
			for _, client := range cl.clients {
				select {
				case client.notifier.Notifier <- c:
				default:
					fmt.Println("buffer is full...")
				}
			}
		}
	}
}

func (cl *ClientSet) UnregisterClient() {
	for {
		cl.mtx.Lock()
		for i, c := range cl.clients {
			select {
			case <-c.notifier.Unregister:
				cl.delete(i)
			default:
				continue
			}
		}
		cl.mtx.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (cl *ClientSet) delete(index int) {
	var tmp []Client
	cl.mtx.Lock()
	for i, c := range cl.clients {
		if i == index {
			continue
		}
		tmp = append(tmp, c)
	}
	cl.clients = tmp
	cl.mtx.Unlock()
}
