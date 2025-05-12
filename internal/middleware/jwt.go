package middleware

import (
	"github.com/google/uuid"
	"github.com/raxaris/ipromise-backend/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/services"
)

// AuthMiddleware – Middleware для проверки Access-токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовка Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Token not found")
			c.Abort()
			return
		}

		// Проверяем, начинается ли заголовок с "Bearer " и содержит ли токен
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		if tokenString == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "No token")
			c.Abort()
			return
		}

		// Валидация токена
		claims, err := services.ValidateAccessToken(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		// Извлекаем user_id
		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization error (user_id)")
			c.Abort()
			return
		}

		// Конвертируем в uuid.UUID
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid user ID")
			c.Abort()
			return
		}

		// Извлекаем роль пользователя
		role, ok := claims["role"].(string)
		if !ok {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization error (role)")
			c.Abort()
			return
		}

		// Передаем user_id и role в контекст Gin
		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next() // Продолжаем выполнение запроса
	}
}
