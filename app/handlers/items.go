package handlers

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"harryd.com/shop/app/models"
)

type ItemHandler struct{
	items []models.Item
}

func NewItemHandler() *ItemHandler {
    items := []models.Item{
        {ID: 1, Name: "A lipstick"},
        {ID: 2, Name: "A pineapple"},
    }
    return &ItemHandler{items: items}
}

func (h *ItemHandler) GetItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"items": h.items})
}

func (h *ItemHandler) GetItem(ctx *gin.Context) {
	itemID := ctx.Param("itemID")
	
	// redundant, may be remove later
	if match, _ := regexp.MatchString("^[0-9]+$", itemID); !match {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	for _, item := range h.items {
        if item.ID == id {
            ctx.JSON(http.StatusOK, item)
            return
        }
    }

    ctx.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}