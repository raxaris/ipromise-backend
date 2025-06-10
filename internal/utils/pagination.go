package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func ParsePaginationParams(c *gin.Context) (limit int, afterCreatedAt *time.Time) {
	limit = 10

	limitParam := c.Query("limit")
	if limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil {
			limit = parsedLimit
		}
	}

	afterParam := c.Query("after")
	if afterParam != "" {
		if parsedTime, err := time.Parse(time.RFC3339, afterParam); err == nil {
			afterCreatedAt = &parsedTime
		}
	}

	return
}

func ParseCursorPaginationParams(c *gin.Context) (limit int, afterCreatedAt *time.Time, afterID *uuid.UUID) {
	limit = 10

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	if t := c.Query("after_created_at"); t != "" {
		if parsedTime, err := time.Parse(time.RFC3339, t); err == nil {
			afterCreatedAt = &parsedTime
		}
	}

	if id := c.Query("after_id"); id != "" {
		if parsedID, err := uuid.Parse(id); err == nil {
			afterID = &parsedID
		}
	}

	return
}
