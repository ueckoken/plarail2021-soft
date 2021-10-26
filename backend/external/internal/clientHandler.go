package internal

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	pb "ueckoken/plarail2021-soft-external/spec"
)

type clientHandler struct {
	upgrader          websocket.Upgrader
	ClientCommand     chan StationState
	ClientChannelSend chan clientChannel
}

type clientChannel struct {
	clientSync chan StationState
	Done       chan bool
}
type clientSendData struct {
	StationName string `json:"station_name"`
	State       string `json:"state"`
}

func (m clientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	var cSync = make(chan StationState, 16)
	var cDone = make(chan bool)
	defer func() { cDone <- true }()
	var cChannel = clientChannel{cSync, cDone}
	m.ClientChannelSend <- cChannel
	go func() {
		r, err := unpackClientSendData(c)
		if err != nil {
			log.Println(err)
			return
		}
		m.ClientCommand <- *r
	}()

	for cChan := range cChannel.clientSync {
		fmt.Println(cChan)
		fmt.Println("sent")
		err := c.WriteJSON(cChan)
		if err != nil {
			fmt.Println("err", err)
			cDone <- true
			close(cDone)
			close(cSync)
			break
		}
	}
}

func unpackClientSendData(c *websocket.Conn) (*StationState, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud clientSendData
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}

	station, ok := pb.Stations_StationId_value[ud.StationName]
	if !ok {
		return nil, fmt.Errorf("bad station format: %s", ud.StationName)
	}

	state, ok := pb.RequestSync_State_value[ud.State]
	if !ok {
		return nil, fmt.Errorf("bad state format: %s", ud.State)
	}
	return &StationState{
		StationID: station,
		State:     state,
	}, nil
}

//go:embed embed/index.html
var IndexHtml []byte

func handleStatic(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(IndexHtml)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
