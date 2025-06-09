package dto

import "mime/multipart"

type MyProfileResponse struct {
	ID          string `json:"id"` // UUID юзера
	Username    string `json:"username"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
	Role        string `json:"role"`
	Followers   int    `json:"followers_count"`
	Following   int    `json:"following_count"`
	Promises    int    `json:"promises_count"`
	BadgesCount int    `json:"badges_count"`
	Bio         string `gorm:"size:160" json:"bio"`
}

type PublicProfileResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AvatarURL   string `json:"avatar_url"`
	Followers   int    `json:"followers_count"`
	Following   int    `json:"following_count"`
	Promises    int    `json:"promises_count"`
	BadgesCount int    `json:"badges_count"`
	Bio         string `gorm:"size:160" json:"bio"`
	IsFollowing bool   `json:"is_following"`
}

type UserLiteResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type UpdateProfileFormRequest struct {
	Username *string               `form:"username"`
	Bio      *string               `form:"bio"`
	Avatar   *multipart.FileHeader `form:"avatar"`
}

type UpdateProfileRequest struct {
	Username  *string `json:"username,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}
