package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dB_DRIVER = "sqlite3"
const dB_Name string = "items.db"
const oPEN_DB_READONLY = dB_Name + "?ro"
const oPEN_DB = dB_Name
const cREATE_DB_SOURCE string = "hack/create-db.sql"

type Page struct {
	Number int
	Size   int
}

func DefaultPage() Page {
	return Page{Number: 1, Size: 100}
}

func Initialize(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	bytes, err := os.ReadFile(cREATE_DB_SOURCE)
	if err != nil {
		log.Fatalf(
			"error occurred when reading the SQL: %s",
			err,
		)
	}

	sql := string(bytes)
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalf(
			"error occurred when creating the database: %s",
			err,
		)
	}
}
