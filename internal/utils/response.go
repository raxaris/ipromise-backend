package utils

import (
	"github.com/gin-gonic/gin"
)

// Унифицированный ответ при ошибке
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"code":    code,
	})
}

// Унифицированный успешный ответ
func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status": "success",
		"data":   data,
		"code":   code,
	})
}

func RespondWithMappedError(c *gin.Context, err error) {
	code := ErrorToStatusCode(err)
	RespondWithError(c, code, err.Error())
}
