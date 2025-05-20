package dto

import "github.com/raxaris/ipromise-backend/internal/models"

type BadgeResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
}

type CreateBadgeRequest struct {
	Code        string `json:"code" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"` // можно потом добавить через MinIO
}

type AssignBadgeRequest struct {
	Username  string `json:"username" binding:"required"`
	BadgeCode string `json:"badge_code" binding:"required"`
}

func MapBadgesToDTO(badges []models.Badge) []BadgeResponse {
	var result []BadgeResponse
	for _, b := range badges {
		result = append(result, BadgeResponse{
			ID:          b.ID.String(),
			Code:        b.Code,
			Title:       b.Title,
			Description: b.Description,
			IconURL:     b.IconURL,
		})
	}
	return result
}
