package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/models"
)

type ItemHandlerImpl struct {
	db *sql.DB
}

func NewItemHandler(db *sql.DB) *ItemHandlerImpl {
	return &ItemHandlerImpl{db: db}
}

func (h *ItemHandlerImpl) GetItems(ctx *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println(err.Error(), " DB ERROR")
		handleDatabaseError(ctx, err)
		return
	}
	defer rows.Close()

	items := []models.Item{}
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			handleDatabaseError(ctx, err)
			return
		}
		items = append(items, item)
	}

	response := models.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    gin.H{"items": items},
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ItemHandlerImpl) GetItem(ctx *gin.Context) {
	itemID := ctx.Param("itemID")

	// redundant, may be remove later
	if match, _ := regexp.MatchString("^[0-9]+$", itemID); !match {
		response := models.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid item_id",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, "invalid item_id")
		return
	}

	row := h.db.QueryRow("SELECT ID, Name FROM items WHERE ID = ?", id)

	var foundItem models.Item
	err = row.Scan(&foundItem.ID, &foundItem.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(ctx, http.StatusNotFound, "item not found")
			return
		}
		handleDatabaseError(ctx, err)
		return
	}

	response := models.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    foundItem,
	}
	ctx.JSON(http.StatusOK, response)
}

func handleDatabaseError(ctx *gin.Context, err error) {
	respondWithError(ctx, http.StatusInternalServerError, err.Error())
}

func respondWithError(ctx *gin.Context, statusCode int, message string) {
	response := models.APIResponse{
		Status:  statusCode,
		Message: message,
	}
	ctx.JSON(statusCode, response)
}
