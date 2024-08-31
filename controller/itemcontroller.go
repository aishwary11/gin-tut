package controller

import (
	"net/http"
	"strconv"

	"gin.com/aishwary11/utils"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Items)
}
func GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, item := range utils.Items {
		if item.ID == id {
			utils.ResponseHelper(c, http.StatusOK, "Item found", item)
			return
		}
	}
	utils.ResponseHelper(c, http.StatusNotFound, "Item not found", nil)
}
