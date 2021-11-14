package internal

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebSocketServer struct {
	upGrader *websocket.Upgrader
}

func (s *WebSocketServer) Start() {
}

func (s *WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := s.upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	//ctx, cancel := context.WithCancel(context.Background())

}
