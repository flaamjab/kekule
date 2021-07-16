package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH string = "items.db"
const dbDriver = "sqlite3"
const openDbReadoly = DB_PATH + "?ro"
const openDb = DB_PATH + "?_fk=true"
const createDbSource string = "hack/create-db.sql"

type Page struct {
	Number int
	Size   int
}

func DefaultPage() Page {
	return Page{Number: 1, Size: 100}
}

func Initialize() {
	db, err := sql.Open(dbDriver, openDb)
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

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("error starting transaction")
	}

	sql = "insert into Item (name, price, category) values (?, ?, ?)"
	stmt, _ := tx.Prepare(sql)

	for ix := 0; ix < 9999; ix++ {
		stmt.Exec(
			fmt.Sprintf("Cringe%d", ix),
			rand.Intn(10000),
			rand.Intn(3)+1,
		)
	}

	tx.Commit()
}
