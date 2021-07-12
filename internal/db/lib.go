package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const createDbSource string = "hack/create-db.sql"

func Initialize(path string) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	bytes, err := os.ReadFile(createDbSource)
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
