package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// do not redeclare DB with := for err
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database...")
	}

	// configure connection pooling
	// set how many connection can be opened simultaneously at most
	// so that we do not keep opening new connections all the time when application runs
	// instead we have a pool of ongoing connections that can be used by different parts of the application
	// and at the same time we make sure this pool is not too big
	DB.SetMaxOpenConns(10)
	// min 5 connection open at all the times
	// if more are needed, can open upto 10 connections
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTables := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTables)
	if err != nil {
		panic("Could not create users table...")
	}

	createEventTables := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventTables)
	if err != nil {
		panic("Could not create events table...")
	}
}
