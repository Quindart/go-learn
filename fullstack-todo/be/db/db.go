package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "my.db")

	if err != nil {
		panic("Could not connected to server")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT 
	)
	`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		fmt.Println("Couldn't create events table", err)
	}
	fmt.Println("Tables created sucessfully!")
}
