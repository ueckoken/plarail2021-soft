package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/clientSync"

	"github.com/gorilla/websocket"
)

type clientHandler struct {
	upgrader          websocket.Upgrader
	ClientCommand     chan clientSync.SingleState
	ClientChannelSend chan clientChannel
}

type clientChannel struct {
	clientSync chan clientSync.ClientSync
	Done       chan bool
}

type data struct {
	Data string `json:"data"`
}

func (m clientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	var cSync = make(chan clientSync.ClientSync, 16)
	var cDone = make(chan bool)
	defer func() { cDone <- true }()
	var cChannel = clientChannel{cSync, cDone}
	m.ClientChannelSend <- cChannel
	for cChan := range cChannel.clientSync {
		fmt.Println(cChan)
		fmt.Println("sent")
		err := c.WriteJSON(cChan)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
}

const page = `
<html>
  <head>
      <title>Hello WebSocket</title>

      <script type="text/javascript">
      var sock = null;
      var data = "";
      function update() {
          var p1 = document.getElementById("plot");
          p1.innerHTML = data;
      };
      window.onload = function() {
          sock = new WebSocket("ws://"+location.host+"/ws");
          sock.onmessage = function(event) {
              var data = JSON.parse(event.data);
              data = data["State"][0].name;
              console.log(data);
              update();
          };
      };
      </script>
  </head>
  <body>
      <div id="header">
          <h1>Hello WebSocket</h1>
      </div>
      <div id="content">
          <div id="plot"></div>
      </div>
  </body>
</html>
`
