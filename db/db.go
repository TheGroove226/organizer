package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//Get DB Func
func Get() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./inventory.db")
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}

// Init Func
func Init() {
	dbc := Get()
	query, err := dbc.Prepare("CREATE TABLE IF NOT EXISTS events(id INTERGER PRIMARY KEY, title TEXT, description TEXT, date TEXT")

	if _, err = query.Exec(); err != nil {
		panic(err)
	}
}
