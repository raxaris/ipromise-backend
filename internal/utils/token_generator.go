package utils

import (
	"github.com/raxaris/ipromise-backend/internal/models"
	"gorm.io/gorm"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raxaris/ipromise-backend/config"
)

func GenerateAccessToken(userID, role string) (string, error) {
	expirationTime := time.Now().Add(1 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func GenerateRefreshToken(userID, role string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidateRefreshTokenFromDB(db *gorm.DB, tokenString string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := db.Where("token = ?", tokenString).First(&refreshToken).Error; err != nil {
		return nil, err
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		db.Delete(&refreshToken)
		return nil, gorm.ErrRecordNotFound
	}

	return &refreshToken, nil
}
