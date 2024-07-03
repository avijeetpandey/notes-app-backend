package main

import (
	"github.com/avijeetpandey/notes-app-backend/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Run(":9000")
}
