package utility

import "github.com/gin-gonic/gin"

// process the error
func ReturnError(statusCode int, message string, context *gin.Context) {
	context.JSON(statusCode, gin.H{
		"message": message,
	})
}

// process the response
func ReturnResponse(statusCode int, message string, context *gin.Context, content any) {
	context.JSON(statusCode, gin.H{
		"message": message,
		"data":    content,
	})
}
