package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func createNote(context *gin.Context) {
	fmt.Println("Add")
}

func getNote(context *gin.Context) {
	fmt.Println("GET")
}

func updateNote(context *gin.Context) {
	fmt.Println("update")
}

func deleteNote(context *gin.Context) {
	fmt.Println("delete")
}
