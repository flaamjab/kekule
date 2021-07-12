package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/api/get/list", getItemList)

	return r
}

func getItemList(c *gin.Context) {
	var r getItemListRequest
	err := c.ShouldBindQuery(&r)
	if err != nil {
		response := errorResponse{
			Name:        "malformed input",
			Description: "check the query string for incorrect values",
		}
		c.JSON(http.StatusBadRequest, response)
	} else {
		fmt.Println(r.HighestPrice)
		fmt.Println(r.LowestPrice)
		fmt.Println(r.Category)

		c.String(http.StatusOK, "There will be an item list, just wait...")
	}
}
