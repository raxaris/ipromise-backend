package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func ParsePaginationParams(c *gin.Context) (limit int, afterCreatedAt *time.Time) {
	// Значение по умолчанию
	limit = 10

	// Получаем значение limit из query-параметров
	limitParam := c.Query("limit")
	if limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil {
			limit = parsedLimit
		}
	}

	// Получаем значение after из query-параметров
	afterParam := c.Query("after")
	if afterParam != "" {
		if parsedTime, err := time.Parse(time.RFC3339, afterParam); err == nil {
			afterCreatedAt = &parsedTime
		}
	}

	return
}
