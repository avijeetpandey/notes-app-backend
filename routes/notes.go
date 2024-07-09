package routes

import (
	"net/http"
	"strconv"

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
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utility.ReturnError(http.StatusBadRequest, constants.PARSE_REQUEST_FAILED, context)
		return
	}

	note, err := models.GetNoteById(noteId)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
	}

	utility.ReturnResponse(http.StatusOK, constants.DONE_MESSAGE, context, note)
}

func updateNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utility.ReturnError(http.StatusBadRequest, constants.PARSE_REQUEST_FAILED, context)
		return
	}

	_, err = models.GetNoteById(noteId)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	var updateNote models.Note
	err = context.ShouldBindJSON(&updateNote)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	err = updateNote.Update()

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	utility.ReturnResponse(http.StatusOK, constants.DONE_MESSAGE, context, updateNote)
}

func deleteNote(context *gin.Context) {
	noteId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utility.ReturnError(http.StatusBadRequest, constants.PARSE_REQUEST_FAILED, context)
		return
	}

	_, err = models.GetNoteById(noteId)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	var updateNote models.Note
	err = context.ShouldBindJSON(&updateNote)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	err = models.Delete(noteId)

	if err != nil {
		utility.ReturnError(http.StatusInternalServerError, constants.SOMETHING_WENT_WRONG_MESSAGE, context)
		return
	}

	utility.ReturnResponse(http.StatusOK, constants.DONE_MESSAGE, context, updateNote)
}
