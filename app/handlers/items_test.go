package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	return db, mock
}

func TestGetItems(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("^SELECT \\* FROM items$").
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "Name"}).
				AddRow(1, "A lipstick").
				AddRow(2, "A pineapple"),
		)

	handler := NewItemHandler(db)
	router := gin.New()
	router.GET("/api/v1/items", handler.GetItems)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Error("Failed to assert 'data' field as a map of string-key and interface-value pairs.")
	}

	items, ok := data["items"].([]interface{})
	assert.True(t, ok, "items must be a slice of interfaces")

	assert.Len(t, items, 2)
}

func TestGetItemValidID(t *testing.T) {
	itemID := 1

	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("^SELECT ID, Name FROM items WHERE ID = ?").
		WithArgs(itemID).
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "Name"}).
				AddRow(itemID, "A lipstick"),
		)

	handler := NewItemHandler(db)
	router := gin.New()
	router.GET("/api/v1/items/:itemID", handler.GetItem)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/items/%d", itemID), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Error("Failed to assert 'data' field as a map of string-key and interface-value pairs.")
	}

	id, ok := data["id"].(float64)
	assert.True(t, ok, "id must be a float64")
	assert.Equal(t, itemID, int(id))

	name, ok := data["name"].(string)
	assert.True(t, ok, "name must be a string")
	assert.Equal(t, "A lipstick", name)
}

func TestGetItemInvalidID(t *testing.T) {
	itemID := 1
	
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("^SELECT ID, Name FROM items WHERE ID = ?").
		WithArgs(itemID).
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "Name"}).
				AddRow(itemID, "A lipstick"),
		)

	handler := NewItemHandler(db)
	router := gin.New()
	router.GET("/api/v1/items/:itemID", handler.GetItem)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/items/abc", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, "invalid item_id", response["message"])
}

func TestGetItemNonExistentID(t *testing.T) {
	itemID := 47

	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("^SELECT ID, Name FROM items WHERE ID = ?").
        WithArgs(itemID).
        WillReturnRows(sqlmock.NewRows([]string{"ID", "Name"}))

	handler := NewItemHandler(db)
	router := gin.New()
	router.GET("/api/v1/items/:itemID", handler.GetItem)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/items/%d", itemID), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, "item not found", response["message"])
}
