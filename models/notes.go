package models

import (
	"time"

	"github.com/avijeetpandey/notes-app-backend/db"
)

// representation of a note
type Note struct {
	ID          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Priority    int    `binding:"required"`
	DateTime    time.Time
}

// save an event
func (n *Note) Save() error {
	query := `
		INSERT INTO NOTES(title,description,priority,dateTime) VALUES (?,?,?,?)
	`

	preparedStatement, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	result, err := preparedStatement.Exec(n.Title, n.Description, n.Priority, n.DateTime)

	if err != nil {
		return nil
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	// update the id of the note created
	n.ID = id

	return err
}
