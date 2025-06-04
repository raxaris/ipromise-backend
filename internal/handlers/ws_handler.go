package handlers

import (
	"github.com/raxaris/ipromise-backend/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"github.com/raxaris/ipromise-backend/internal/ws"
)

type WebSocketHandler struct {
	hub                 *ws.NotificationHub
	notificationService services.NotificationService
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketHandler(hub *ws.NotificationHub, notificationService services.NotificationService) *WebSocketHandler {
	return &WebSocketHandler{hub: hub, notificationService: notificationService}
}

// ServeWS godoc
// @Summary WebSocket подключение для получения уведомлений
// @Description Устанавливает WebSocket-соединение и передаёт оффлайн-уведомления, если есть
// @Tags websocket
// @Security BearerAuth
// @Produce json
// @Success 200 {string} string "WebSocket connection established"
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /ws/notifications [get]
func (h *WebSocketHandler) ServeWS(c *gin.Context) {
	log.Println("🚀 WebSocket handler triggered")
	userID, _, err := utils.ExtractUserFromRequest(c)
	if err != nil {
		utils.RespondWithMappedError(c, utils.ErrUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("❌ WebSocket upgrade failed: %v", err)
		utils.RespondWithMappedError(c, utils.ErrInternal)
		return
	}

	client := &ws.Client{
		UserID: userID.String(),
		Conn:   conn,
		Send:   make(chan interface{}, 256),
		Hub:    h.hub,
	}

	h.hub.Register(client)
	client.Send <- map[string]string{
		"type":    "hello",
		"message": "WebSocket connected successfully 🎉",
	}
	go h.notificationService.DeliverOfflineNotifications(userID)
	go client.ReadPump()
	go client.WritePump()
}
