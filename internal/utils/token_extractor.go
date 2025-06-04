package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/services"
	"strings"
)

func ExtractUserFromRequest(c *gin.Context) (uuid.UUID, string, error) {
	var tokenString string

	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenString = parts[1]
		}
	}

	if tokenString == "" {
		cookieToken, err := c.Cookie("access_token")
		if err == nil && cookieToken != "" {
			tokenString = cookieToken
		}
	}

	if tokenString == "" {
		return uuid.Nil, "", errors.New("access token not provided")
	}

	claims, err := services.ValidateAccessToken(tokenString)
	if err != nil {
		return uuid.Nil, "", err
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, "", errors.New("user_id not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, "", err
	}

	role, ok := claims["role"].(string)
	if !ok {
		return uuid.Nil, "", errors.New("role not found")
	}

	return userID, role, nil
}
