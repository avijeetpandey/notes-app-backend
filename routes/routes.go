package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// notes route
	server.POST("/notes", createNote)
	server.GET("/note/:id", getNote)
	server.PUT("/note/:id", updateNote)
	server.DELETE("/note/:id", deleteNote)

	// health route
	server.GET("/health", health)
}
