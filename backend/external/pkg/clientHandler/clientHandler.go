package clientHandler

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-external/pkg/servo"
	"ueckoken/plarail2021-soft-external/pkg/syncController"
	pb "ueckoken/plarail2021-soft-external/spec"

	"github.com/gorilla/websocket"
)

type ClientHandler struct {
	upgrader          websocket.Upgrader
	ClientCommand     chan syncController.StationState
	ClientChannelSend chan ClientChannel
}

type ClientChannel struct {
	ClientSync chan syncController.StationState
	Done       chan struct{}
}
type clientSendData struct {
	StationName string `json:"station_name"`
	State       string `json:"state"`
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("responsing")
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	ctx, cancel := context.WithCancel(context.Background())
	var cSync = make(chan syncController.StationState, 16)
	var cDone = make(chan struct{})
	var cChannel = ClientChannel{cSync, cDone}
	m.ClientChannelSend <- cChannel
	fmt.Println("added")
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(20 * time.Second))
		return nil
	})
	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientCommand(ctx, c, &m)
	go handleClientPing(ctx, c)
	for cChan := range cChannel.ClientSync {
		fmt.Println(cChan)
		fmt.Println("sent")
		dat := clientSendData{StationName: pb.Stations_StationId_name[cChan.StationID], State: pb.RequestSync_State_name[cChan.State]}
		err := c.WriteJSON(dat)
		if err != nil {
			fmt.Println("err", err)
			cDone <- struct{}{}
			cancel()
			break
		}
	}
}

func handleClientPing(ctx context.Context, c *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(1*time.Second)); err != nil {
				fmt.Println("ping:", err)
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func handleClientCommand(ctx context.Context, c *websocket.Conn, m *ClientHandler) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			r, err := unpackClientSendData(c)
			if err != nil {
				log.Println(err)
				return
			}
			m.ClientCommand <- *r
		}
	}
}

func unpackClientSendData(c *websocket.Conn) (*syncController.StationState, error) {
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
	fmt.Printf("Received: StationID:%d, State:%d\n", station, state)
	return &syncController.StationState{
		StationState: servo.StationState{
			StationID: station,
			State:     state,
		},
	}, nil
}

//go:embed embed/index.html
var IndexHtml []byte

func HandleStatic(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(IndexHtml)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
