package ws

import (
	"log"
	"sync"

	"github.com/raxaris/ipromise-backend/internal/ws/types"
)

type NotificationHub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	send       chan types.MessageEnvelope
	mu         sync.RWMutex
}

func NewNotificationHub() *NotificationHub {
	return &NotificationHub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		send:       make(chan types.MessageEnvelope, 256),
	}
}

func (h *NotificationHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.UserID] = client
			h.mu.Unlock()
			log.Printf("🟢 User %s connected", client.UserID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				delete(h.clients, client.UserID)
				close(client.Send)
				log.Printf("🔴 User %s disconnected", client.UserID)
			}
			h.mu.Unlock()

		case msg := <-h.send:
			h.mu.RLock()
			client, ok := h.clients[msg.ToUserID]
			h.mu.RUnlock()
			if ok {
				select {
				case client.Send <- msg.Data:
				default:
					log.Printf("⚠️ Send channel full for user %s", msg.ToUserID)
				}
			}
		}
	}
}

func (h *NotificationHub) IsUserConnected(userID string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, ok := h.clients[userID]
	return ok
}

func (h *NotificationHub) Register(c *Client) {
	h.register <- c
}

func (h *NotificationHub) Unregister(c *Client) {
	h.unregister <- c
}

func (h *NotificationHub) SendToUser(toUserID string, payload interface{}) {
	h.send <- types.MessageEnvelope{
		ToUserID: toUserID,
		Data:     payload,
	}
}
