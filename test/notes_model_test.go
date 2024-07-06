package test

import (
	"testing"

	"github.com/avijeetpandey/notes-app-backend/models"
)

func TestNoteCreationFromModel(t *testing.T) {
	var note models.Note = models.Note{
		ID:          2322,
		Title:       "test note",
		Description: "description of note",
		DateTime:    "",
	}

	if note.Title != "test note" {
		t.Fatalf("unable to assign title properly")
	}
}
