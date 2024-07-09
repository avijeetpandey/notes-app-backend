package models

import (
	"fmt"

	"github.com/avijeetpandey/notes-app-backend/db"
)

// representation of a note
type Note struct {
	ID          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Priority    int
}

// save an event
func (n *Note) Save() error {
	query := `
		INSERT INTO NOTES(title,description,priority) VALUES (?,?,?)
	`

	preparedStatement, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	result, err := preparedStatement.Exec(n.Title, n.Description, n.Priority)

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

// get all notes
func GetAllNotes() ([]Note, error) {
	query := `SELECT * FROM NOTES`
	row, err := db.GlobalDB.Query(query)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var notes []Note

	for row.Next() {
		var note Note
		err := row.Scan(&note.ID, &note.Title, &note.Description, &note.Priority)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

// get a note by ID
func GetNoteById(id int64) (*Note, error) {
	query := fmt.Sprintf("SELECT * FROM NOTES WHERE id= %d", id)
	row := db.GlobalDB.QueryRow(query)

	var note Note
	err := row.Scan(&note.ID, &note.Title, &note.Description, &note.Priority)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

// update a note
func (n Note) Update() error {
	query := `
		UPDATE NOTES 
		SET title = ?, description = ?, priority = ?
		WHERE id = ?	
	`

	preparedStatement, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	_, err = preparedStatement.Exec(&n.ID, &n.Title, &n.Description, &n.Priority)

	if err != nil {
		return err
	}

	return nil
}

// delete a note
func Delete(id int64) error {
	query := `DELETE from NOTES WHERE id = ?`
	preparedStatement, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	_, err = preparedStatement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
