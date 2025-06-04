package ws

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	UserID string
	Conn   *websocket.Conn
	Send   chan interface{}
	Hub    *NotificationHub
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Printf("⚠️ Error closing connection (readPump): %v", err)
		}
	}()

	c.Conn.SetReadLimit(maxMessageSize)

	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("⚠️ SetReadDeadline error (initial): %v", err)
		return
	}

	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Printf("⚠️ SetReadDeadline error (pong): %v", err)
			return err
		}
		return nil
	})

	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("⚠️ Unexpected close error: %v", err)
			} else {
				log.Printf("ℹ️ Connection closed (readPump): %v", err)
			}
			break
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.Conn.Close(); err != nil {
			log.Printf("⚠️ Error closing connection (writePump): %v", err)
		}
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("⚠️ SetWriteDeadline error: %v", err)
				return
			}

			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("⚠️ Error sending close message: %v", err)
				}
				return
			}

			if err := c.Conn.WriteJSON(message); err != nil {
				log.Printf("❌ WriteJSON error: %v", err)
				return
			}

		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("⚠️ SetWriteDeadline error (ping): %v", err)
				return
			}

			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("⚠️ Ping error: %v", err)
				return
			}
		}
	}
}
