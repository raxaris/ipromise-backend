package mappers

import (
	"github.com/raxaris/ipromise-backend/internal/dto"
	"github.com/raxaris/ipromise-backend/internal/models"
	"github.com/raxaris/ipromise-backend/internal/ws/types"
)

type NotificationMapper struct{}

func NewNotificationMapper() *NotificationMapper {
	return &NotificationMapper{}
}

func (m *NotificationMapper) ToDTO(n *models.Notification) dto.NotificationResponse {
	return dto.NotificationResponse{
		ID:      n.ID.String(),
		Type:    n.Type,
		Message: n.Message,
		RelatedID: func() string {
			if n.RelatedID != nil {
				return n.RelatedID.String()
			}
			return ""
		}(),
		IsRead:    n.IsRead,
		CreatedAt: n.CreatedAt,
	}
}

func (m *NotificationMapper) ToPayload(n *models.Notification) types.MessagePayload {
	return types.MessagePayload{
		Type:    "notification",
		Payload: m.ToDTO(n),
	}
}

func (m *NotificationMapper) ToListDTO(list []models.Notification) []dto.NotificationResponse {
	result := make([]dto.NotificationResponse, 0, len(list))
	for _, n := range list {
		result = append(result, m.ToDTO(&n))
	}
	return result
}
