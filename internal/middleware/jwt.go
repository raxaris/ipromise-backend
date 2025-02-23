package middleware

import (
	"github.com/google/uuid"
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
			c.Abort()
			return
		}

		// Проверяем, начинается ли заголовок с "Bearer " и содержит ли токен
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
			c.Abort()
			return
		}

		// Валидация токена
		claims, err := services.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		// Извлекаем user_id
		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации (user_id)"})
			c.Abort()
			return
		}

		// Конвертируем в uuid.UUID
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат ID пользователя"})
			c.Abort()
			return
		}

		// Извлекаем роль пользователя
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка авторизации (role)"})
			c.Abort()
			return
		}

		// Передаем user_id и role в контекст Gin
		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next() // Продолжаем выполнение запроса
	}
}
