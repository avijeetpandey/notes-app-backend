package routes

import (
	"fmt"
	"net/http"

	"github.com/avijeetpandey/notes-app-backend/constants"
	"github.com/avijeetpandey/notes-app-backend/models"
	"github.com/avijeetpandey/notes-app-backend/utility"
	"github.com/gin-gonic/gin"
)

// function to create note
func createNote(context *gin.Context) {
	var note models.Note
	err := context.ShouldBindJSON(&note)

	if err != nil {
		utility.ReturnError(http.StatusBadRequest, constants.PARSE_REQUEST_FAILED, context)
		return
	}

	err = note.Save()

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.UNABLE_TO_CREATE_RESOURCE, context)
		return
	}

	// return the created note
	utility.ReturnResponse(http.StatusCreated, constants.DONE_MESSAGE, context, note)
}

func getAllNotes(context *gin.Context) {
	notes, err := models.GetAllNotes()

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	utility.ReturnResponse(http.StatusOK, constants.DONE_MESSAGE, context, notes)
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
