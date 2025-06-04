package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/raxaris/ipromise-backend/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := utils.ExtractUserFromRequest(c)
		if err != nil {
			utils.RespondWithMappedError(c, err)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}
}
