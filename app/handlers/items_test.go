package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetItems(t *testing.T) {
	handler := NewItemHandler()
	router := gin.Default()
	router.GET("/items", handler.GetItems)

	req, err := http.NewRequest(http.MethodGet, "/items", nil)
	assert.NoError(t, err)
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	items, ok := response["items"].([]interface{})
	assert.True(t, ok, "items must be a slice of interfaces")

	assert.Len(t, items, len(handler.items))
}

func TestGetItemValidID(t *testing.T) {
	handler := NewItemHandler()
	router := gin.Default()
	router.GET("/items/:itemID", handler.GetItem)

	req, err := http.NewRequest(http.MethodGet, "/items/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	id, ok := response["id"].(float64)
	assert.True(t, ok, "id must be a float64")
	assert.Equal(t, 1, int(id))
}

func TestGetItemInvalidID(t *testing.T) {
	handler := &ItemHandler{}
	router := gin.Default()
	router.GET("/items/:itemID", handler.GetItem)

	req, err := http.NewRequest(http.MethodGet, "/items/abc", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, "invalid item_id", response["error"])
}