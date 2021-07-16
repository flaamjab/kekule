package api

import (
	"fmt"
	"net/http"

	"github.com/flaamjab/kekule/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
)

func router() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")

	itemGroup := apiGroup.Group("/item")
	itemGroup.GET("", getItem)
	itemGroup.GET("/list", getItemList)
	itemGroup.POST("", postItem)
	itemGroup.PUT("", putItem)
	itemGroup.DELETE("", deleteItem)

	apiGroup.GET("/category", getCategory)

	return r
}

func getItem(c *gin.Context) {
	var r getItemRequest
	err := c.ShouldBindQuery(&r)
	if err == nil {
		item, err := db.GetItem(r.Id)
		if err == nil {
			if item != nil {
				r := getItemResponse{
					Result: "success",
					Item:   *item,
				}
				c.JSON(http.StatusOK, r)
			} else {
				r := resultResponse{
					Result: "failure",
					Description: fmt.Sprintf(
						"item with ID %d does not exist",
						r.Id,
					),
				}
				c.JSON(http.StatusNotFound, r)
			}
		} else {
			fmt.Println(err)
			r := resultResponse{
				Result:      "server error",
				Description: "an error occurred when fetching the item",
			}
			c.JSON(http.StatusInternalServerError, r)
		}
	} else {
		r := resultResponse{
			Result:      "bad request",
			Description: "make sure the provided ID is a positive number",
		}
		c.JSON(http.StatusBadRequest, r)
	}
}

func getItemList(c *gin.Context) {
	page := 1
	limit := 100
	r := getItemListRequestFilters{
		Page:  &page,
		Limit: &limit,
	}

	if err := c.ShouldBindQuery(&r); err == nil {
		page := db.Page{Number: *r.Page, Size: *r.Limit}
		filters := db.ItemFilters{
			Category:   r.Category,
			LowerPrice: r.LowerPrice,
			UpperPrice: r.UpperPrice,
		}

		items, err := db.GetItemList(page, filters)
		if err == nil {
			r := getItemListResponse{Result: "success", Items: items}
			c.JSON(http.StatusOK, r)
		} else {
			r := resultResponse{
				Result:      "server error",
				Description: "error fetching item list",
			}
			c.JSON(http.StatusInternalServerError, r)
		}
	} else {
		r := resultResponse{
			Result:      "bad request",
			Description: "check the query string for incorrect values",
		}
		c.JSON(http.StatusBadRequest, r)
	}
}

func postItem(c *gin.Context) {
	var r postItemRequest
	if err := c.ShouldBindJSON(&r); err == nil {
		id, err := db.NewItem(r.Name, r.Price, r.Category)
		if err == nil {
			r := newItemResponse{Result: "success", Id: id}
			c.JSON(http.StatusOK, r)
		} else {
			err := err.(sqlite3.Error)
			switch err.Code {
			case sqlite3.ErrConstraint:
				r := resultResponse{
					Result:      "forbidden",
					Description: "this action is not permitted",
				}
				c.JSON(http.StatusForbidden, r)
			default:
				r := resultResponse{
					Result:      "server error",
					Description: "failed to create an item",
				}
				c.JSON(http.StatusInternalServerError, r)
			}

		}
	} else {
		r := resultResponse{
			Result: "bad request",
			Description: "make sure that all required input parameters " +
				"are specified and have correct values",
		}
		c.JSON(http.StatusBadRequest, r)
	}
}

func putItem(c *gin.Context) {
	var r putItemRequest
	if err := c.ShouldBindJSON(&r); err == nil {
		item, err := r.ToItem()
		if err == nil {
			if item != nil {
				err := db.UpdateItem(*item)
				if err == nil {
					r := resultResponse{
						Result:      "success",
						Description: "item updated",
					}
					c.JSON(http.StatusOK, r)
				} else {
					fmt.Println(err)
					r := resultResponse{
						Result:      "server error",
						Description: "error updating the item",
					}
					c.JSON(http.StatusInternalServerError, r)
				}
			} else {
				r := resultResponse{
					Result: "not found",
					Description: fmt.Sprintf(
						"item with ID %d does not exist",
						r.Id,
					),
				}
				c.JSON(http.StatusNotFound, r)
			}
		} else {
			err := err.(sqlite3.Error)
			switch err.Code {
			case sqlite3.ErrConstraint:
				r := resultResponse{
					Result:      "forbidden",
					Description: "this action is not permitted",
				}
				c.JSON(http.StatusForbidden, r)
			default:
				r := resultResponse{
					Result:      "server error",
					Description: "error fetching the item",
				}
				c.JSON(http.StatusInternalServerError, r)
			}
		}
	} else {
		r := resultResponse{
			Result: "bad request",
			Description: "make sure that ID is provided " +
				"and all values have correct type",
		}
		c.JSON(http.StatusBadRequest, r)
	}
}

func deleteItem(c *gin.Context) {
	var r deleteItemRequest
	if err := c.ShouldBindJSON(&r); err == nil {
		deleted, err := db.DeleteItem(r.Id)
		if err == nil {
			if deleted {
				r := resultResponse{
					Result:      "success",
					Description: "item deleted",
				}
				c.JSON(http.StatusOK, r)
			} else {
				r := resultResponse{
					Result: "failure",
					Description: fmt.Sprintf(
						"item with ID %d does not exist",
						r.Id,
					),
				}
				c.JSON(http.StatusBadRequest, r)
			}
		} else {
			r := resultResponse{
				Result:      "server error",
				Description: "error deleting the item",
			}
			c.JSON(http.StatusInternalServerError, r)
		}
	}
}

func getCategory(c *gin.Context) {
	c.String(http.StatusOK, "Category would be returned...")
}
