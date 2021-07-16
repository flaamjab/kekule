package api

import "github.com/flaamjab/kekule/internal/db"

type resultResponse struct {
	Result      string `json:"result"`
	Description string `json:"description"`
}

type newItemResponse struct {
	Result string `json:"result"`
	Id     int64  `json:"id"`
}

type getItemResponse struct {
	Result string  `json:"result"`
	Item   db.Item `json:"item"`
}

type getItemListResponse struct {
	Result string    `json:"result"`
	Items  []db.Item `json:"items"`
}

type getCategoryResponse struct {
	Result   string `json:"result"`
	Category int    `json:"category"`
}
