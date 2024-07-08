package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalDB *sql.DB

const DATABASE_TYPE = "sqlite3"
const DATABASE_NAME = "notes.db"

func InitDB() {
	conn, err := sql.Open(DATABASE_TYPE, DATABASE_NAME)

	GlobalDB = conn

	if err != nil {
		panic("could not connect to the database")
	}

	fmt.Println("successfully connected to the database")

	GlobalDB.SetMaxOpenConns(10)
	GlobalDB.SetMaxIdleConns(5)

	// create tables after connecting to the database
	createTables()
}

func createTables() {
	// create notes table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	  email TEXT NOT NULL UNIQUE,
	  password TEXT NOT NULL 
	)
  `

	_, err := GlobalDB.Exec(createUsersTable)

	if err != nil {
		panic("unable to create users table")
	}

	createNotesTable := `
  CREATE TABLE IF NOT EXISTS notes (
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  title TEXT NOT NULL,
	  description TEXT NOT NULL,
	  priority INTEGER NOT NULL
  )
  `

	_, err = GlobalDB.Exec(createNotesTable)

	if err != nil {
		panic("could not create notes table")
	}

	fmt.Println("TABLES CREATED SUCCESSFULLY")
}
