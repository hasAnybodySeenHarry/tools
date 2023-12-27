package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	api "harryd.com/tools/app/models/api"
	domain "harryd.com/tools/app/models/domain"
)

type ItemHandlerImpl struct {
	db *sql.DB
}

func validatePayload(actual map[string]interface{}, expected map[string]bool) bool {
	for key := range actual {
		fmt.Println("KEY ", key)
		if _, exists := expected[key]; !exists {
			return true
		}
	}
	return false
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

	items := []domain.Item{}
	for rows.Next() {
		var item domain.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			handleDatabaseError(ctx, err)
			return
		}
		items = append(items, item)
	}

	response := api.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    gin.H{"items": items},
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ItemHandlerImpl) GetItem(ctx *gin.Context) {
	isValid, id := validateItemID(ctx.Param("itemID"))
	if !isValid {
		respondWithError(ctx, http.StatusBadRequest, "invalid item_id")
		return
	}

	row := h.db.QueryRow("SELECT ID, Name FROM items WHERE ID = ?", id)

	var foundItem domain.Item
	err := row.Scan(&foundItem.ID, &foundItem.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(ctx, http.StatusNotFound, "item not found")
			return
		}
		handleDatabaseError(ctx, err)
		return
	}

	response := api.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    foundItem,
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ItemHandlerImpl) CreateItem(ctx *gin.Context) {
	var request api.ItemCreateRequest

	if err := ctx.BindJSON(&request); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.db.Exec("INSERT INTO items (Name) VALUES (?)", request.Name)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	response := api.APIResponse{
		Status:  http.StatusCreated,
		Message: "success",
		Data: domain.Item{
			ID:   int(insertedID),
			Name: request.Name,
		},
	}
	ctx.JSON(http.StatusCreated, response)
}

func (h *ItemHandlerImpl) DeleteItem(ctx *gin.Context) {
	isValid, id := validateItemID(ctx.Param("itemID"))
	if !isValid {
		respondWithError(ctx, http.StatusBadRequest, "invalid item_id")
		return
	}

	result, err := h.db.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		respondWithError(ctx, http.StatusNotFound, "item not found")
		return
	}

	response := api.APIResponse{
		Status:  http.StatusNoContent,
		Message: "success",
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *ItemHandlerImpl) UpdateItem(ctx *gin.Context) {
	isValid, id := validateItemID(ctx.Param("itemID"))
	if !isValid {
		respondWithError(ctx, http.StatusBadRequest, "invalid item_id")
		return
	}

	var request api.ItemUpdateRequest

	if err := ctx.BindJSON(&request); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// might remove this preflight later
	var existingItem domain.Item
	err := h.db.QueryRow("SELECT ID, Name FROM items WHERE ID = ?", id).Scan(&existingItem.ID, &existingItem.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(ctx, http.StatusNotFound, "item not found")
			return
		}
		handleDatabaseError(ctx, err)
		return
	}

	result, err := h.db.Exec("UPDATE items SET Name = ? WHERE ID = ?", request.Name, id)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	// without preflight, using mysql, we need to compare to see
	// if the data were same or there was no data existed if rows
	// affected was 0.
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(ctx, http.StatusOK, "data is kept")
		return
	}

	updatedItem := domain.Item{}
	err = h.db.QueryRow("SELECT ID, Name FROM items WHERE ID = ?", id).Scan(&updatedItem.ID, &updatedItem.Name)
	if err != nil {
		handleDatabaseError(ctx, err)
		return
	}

	response := api.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    updatedItem,
	}
	ctx.JSON(http.StatusOK, response)
}

func validateItemID(itemID string) (isValid bool, id int) {
	// redundant, may be remove later
	if match, _ := regexp.MatchString("^[0-9]+$", itemID); !match {
		return false, 0
	}

	id, err := strconv.Atoi(itemID)
	if err != nil {
		return false, 0
	}

	if id == 0 {
		return false, 0
	}

	return true, id
}

func handleDatabaseError(ctx *gin.Context, err error) {
	respondWithError(ctx, http.StatusInternalServerError, err.Error())
}

func respondWithError(ctx *gin.Context, statusCode int, message string) {
	response := api.APIResponse{
		Status:  statusCode,
		Message: message,
	}
	ctx.JSON(statusCode, response)
}
