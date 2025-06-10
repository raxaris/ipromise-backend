package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"user not found"`
	Code    int    `json:"code" example:"404"`
}

type SuccessResponse struct {
	Status string      `json:"status" example:"success"`
	Code   int         `json:"code" example:"200"`
	Data   interface{} `json:"data"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"code":    code,
	})
}

func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	var wrapped interface{}

	if msg, ok := data.(string); ok {
		wrapped = gin.H{"message": msg}
	} else {
		wrapped = data
	}

	c.JSON(code, gin.H{
		"status": "success",
		"code":   code,
		"data":   wrapped,
	})
}

func RespondWithMappedError(c *gin.Context, err error) {
	code := ErrorToStatusCode(err)
	RespondWithError(c, code, err.Error())
}
