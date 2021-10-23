package internal

import (
	"log"
	"net/http"
  "ueckoken/plarail2021-soft-external/pkg/clientSync"
	"github.com/gorilla/websocket"
)

type clientHandler struct{
  upgrader websocket.Upgrader
  ClientCommand chan clientSync.SingleState
  ClientChannelSend chan clientChannel
}

type clientChannel struct{
  clientSync chan clientSync.ClientSync
  Done chan bool
}

func (m clientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
  c, err := m.upgrader.Upgrade(w,r,nil)
  if err != nil{
    log.Println(err)
    return
  }
  defer c.Close()
  var cSync = make(chan clientSync.ClientSync, 16)
  var cDone = make(chan bool)
  var cChannel = clientChannel{cSync, cDone}
  m.ClientChannelSend<-cChannel
  for {

  }
}
