package api

type getItemRequest struct {
}

type getItemListRequest struct {
	LowestPrice  float64 `form:"lowest_price"`
	HighestPrice float64 `form:"highest_price"`
	Category     int     `form:"category"`
}
