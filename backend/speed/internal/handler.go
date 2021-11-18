package internal

import (
	"net/http"
	"ueckoken/plarail2021-soft-speed/pkg/storeSpeed"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	n := make(chan storeSpeed.TrainConf)
	unregister := make(chan struct{})
	notifier := Client{ClientNotifier{n, unregister}}
	m.ClientNotification <- notifier
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(20 * time.Second))
		return nil
	})
	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientPing(ctx, c)
	go handleClientCommand(ctx, c, &m)
	for notification := range notifier.notifier.Notifier {
		err := c.WriteJSON(notification)
		if err != nil {
			notifier.notifier.Unregister <- struct{}{}
			cancel()
			break
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

func unpackClientSendData(c *websocket.Conn) (*storeSpeed.TrainConf, error) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("websocket read failed: %e", err)
	}
	var ud storeSpeed.TrainConf
	err = json.Unmarshal(msg, &ud)
	if err != nil {
		return nil, fmt.Errorf("bad json format: %e", err)
	}
	return &ud, nil
}
