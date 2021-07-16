package api

import "github.com/flaamjab/kekule/internal/db"

type getItemRequest struct {
	Id int64 `form:"id"`
}

type getItemListRequestFilters struct {
	Page       *int     `form:"page" binding:"min=1"`
	Limit      *int     `form:"limit" binding:"max=100"`
	LowerPrice *float64 `form:"lowest_price"`
	UpperPrice *float64 `form:"highest_price"`
	Category   *int64   `form:"category"`
}

type postItemRequest struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=0"`
	Category int64   `json:"category" binding:"required"`
}

type putItemRequest struct {
	Id       int64    `form:"id"`
	Name     *string  `form:"name"`
	Price    *float64 `form:"price" binding:"omitempty,min=0"`
	Category *int64   `form:"category"`
}

func (r putItemRequest) ToItem() (*db.Item, error) {
	item, err := db.GetItem(r.Id)
	if err != nil {
		return nil, err
	}

	if r.Name != nil {
		item.Name = *r.Name
	}

	if r.Price != nil {
		item.Price = *r.Price
	}

	if r.Category != nil {
		item.Category = *r.Category
	}

	return item, nil
}

type deleteItemRequest struct {
	Id int64 `form:"id" binding:"required"`
}
