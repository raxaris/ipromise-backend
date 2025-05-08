package dto

import "github.com/raxaris/ipromise-backend/internal/models"

type BadgeResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
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
