package dto

type UpdateUserRequest struct {
	Username  *string `json:"username" binding:"omitempty,min=3,max=30"`
	Role      *string `json:"role,omitempty"`       // Только для админа
	AvatarURL *string `json:"avatar_url,omitempty"` // Новый аватар
}

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	AvatarURL string `json:"avatar_url"`
}
