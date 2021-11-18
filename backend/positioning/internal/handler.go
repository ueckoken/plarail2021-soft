package internal

import (
	"net/http"
	"ueckoken/plarail2021-soft-positioning/pkg/addressChecker"
	"ueckoken/plarail2021-soft-positioning/pkg/trainState"

	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type ClientHandler struct {
	upgrader           websocket.Upgrader
	ClientNotification chan ClientNotifier
	Checker            addressChecker.AddressChecker
}

type ClientNotifier struct {
	Notifier   chan trainState.PositionAndSpeed
	Unregister chan struct{}
}

func (m ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()

	ctx, cancel := context.WithCancel(context.Background())
	n := make(chan trainState.PositionAndSpeed)
	unregister := make(chan struct{})
	notifier := ClientNotifier{n, unregister}
	m.ClientNotification <- notifier
	c.SetPongHandler(func(string) error {
		c.SetReadDeadline(time.Now().Add(20 * time.Second))
		cancel()
		return nil
	})
	c.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed")
		cancel()
		return nil
	})
	go handleClientPing(ctx, c)
	for notification := range notifier.Notifier {
		err := c.WriteJSON(notification)
		if err != nil {
			notifier.Unregister <- struct{}{}
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
