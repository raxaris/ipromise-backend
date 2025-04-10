package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserIDFromContext извлекает user_id из gin.Context и приводит к uuid.UUID
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, errors.New("user_id не найден в контексте")
	}

	id, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("user_id неверного формата")
	}

	return id, nil
}
