package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Item struct {
	Id       int64   `json:"id"`
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category int64   `json:"category"`
}

type ItemFilters struct {
	LowerPrice *float64
	UpperPrice *float64
	Category   *int64
}

func (itf ItemFilters) toString() string {
	filters := make([]string, 0)

	if itf.Category != nil {
		filters = append(filters, fmt.Sprintf("category = %d", *itf.Category))
	}

	if itf.LowerPrice != nil {
		filters = append(filters, fmt.Sprintf("price >= %f", *itf.LowerPrice))
	}

	if itf.UpperPrice != nil {
		filters = append(filters, fmt.Sprintf("price < %f", *itf.UpperPrice))
	}

	if len(filters) > 0 {
		return "where " + strings.Join(filters, " and ")
	}

	return ""
}

func NewItem(
	name string,
	price float64,
	category int64,
) (int64, error) {
	db, err := sql.Open(dB_DRIVER, oPEN_DB)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sql := "insert into Item (name, price, category) values (?, ?, ?)"
	r, err := db.Exec(sql, name, price, category)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Fatalf("the database or driver does not support LastInsertId")
	}

	return id, nil
}

func GetItem(id int64) (*Item, error) {
	db, err := sql.Open(dB_DRIVER, oPEN_DB_READONLY)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sql := "select * from Item where id = ?"
	rows, err := db.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemFound := rows.Next()
	if !itemFound {
		return nil, nil
	}

	var item Item
	err = rows.Scan(
		&item.Id,
		&item.Sku,
		&item.Name,
		&item.Price,
		&item.Category,
	)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetItemList(
	page Page,
	filters ItemFilters,
) ([]Item, error) {
	db, err := sql.Open(dB_DRIVER, oPEN_DB_READONLY)
	if err != nil {
		return nil, err
	}

	startRow := (page.Number-1)*page.Size + 1
	endRow := startRow + page.Size

	sql := fmt.Sprintf(
		`select * from (
			select row_number() over (order by id) as row_num, *
			from Item
			%s
		) as row_constrained_result
		where row_num >= %d and row_num < %d`,
		filters.toString(), startRow, endRow,
	)

	fmt.Println(sql)

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0)
	for rows.Next() {
		rowNum := 0
		item := Item{}
		rows.Scan(
			&rowNum,
			&item.Id,
			&item.Sku,
			&item.Name,
			&item.Price,
			&item.Category,
		)
		items = append(items, item)
	}

	return items, nil
}

func UpdateItem(item Item) error {
	db, err := sql.Open(dB_DRIVER, oPEN_DB)
	if err != nil {
		return err
	}

	sql := "update Item set name = ?, price = ?, category = ? where id = ?"
	_, err = db.Exec(sql, item.Name, item.Price, item.Category, item.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteItem(id int64) (bool, error) {
	db, err := sql.Open(dB_DRIVER, oPEN_DB)
	if err != nil {
		return false, err
	}
	defer db.Close()

	sql := "delete from Item where id = ?"
	r, err := db.Exec(sql, id)
	if err != nil {
		return false, err
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Fatalln("the database or driver does support RowsAffected")
	}

	if n > 0 {
		return true, nil
	}

	return false, nil
}
