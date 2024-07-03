package main

import (
	"github.com/avijeetpandey/notes-app-backend/db"
	"github.com/avijeetpandey/notes-app-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":9000")
}
