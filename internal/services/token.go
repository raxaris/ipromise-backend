package services

import (
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raxaris/ipromise-backend/config"
)

// GenerateAccessToken – создает Access-токен
func GenerateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

// GenerateRefreshToken – создает Refresh-токен
func GenerateRefreshToken(userID string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

// ValidateAccessToken – проверяет Access-токен
func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

// ValidateRefreshTokenFromDB – проверяет `refresh_token` в БД
func ValidateRefreshTokenFromDB(db *gorm.DB, tokenString string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := db.Where("token = ?", tokenString).First(&refreshToken).Error; err != nil {
		return nil, err
	}

	// Проверяем, не истёк ли токен
	if time.Now().After(refreshToken.ExpiresAt) {
		db.Delete(&refreshToken) // Удаляем истёкший токен
		return nil, gorm.ErrRecordNotFound
	}

	return &refreshToken, nil
}
