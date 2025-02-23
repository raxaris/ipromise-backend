package dto

type UpdateUserRequest struct {
	Username *string `json:"username,omitempty"`
	Role     *string `json:"role,omitempty"` // Только для админов
}
