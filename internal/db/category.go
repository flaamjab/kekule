package db

import (
	"database/sql"
)

type Category struct {
	Id   int
	Name string
}

func GetCategory(id int64) (*Category, error) {
	db, err := sql.Open(dbDriver, openDbReadoly)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := "select * from Category where id = ?"
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var c Category
	if rows.Next() {
		rows.Scan(&c.Id, &c.Name)
	} else if rows.Err() != nil {
		return nil, rows.Err()
	} else {
		return nil, nil
	}

	return &c, nil
}
