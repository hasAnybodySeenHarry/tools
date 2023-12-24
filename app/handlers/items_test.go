package handlers

import "testing"

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

func TestGetItems(t *testing.T) {
	// handler := NewItemHandler()
	// router := gin.New()
	// router.GET("/items", handler.GetItems)

	// req, err := http.NewRequest(http.MethodGet, "/items", nil)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

	// var response map[string]interface{}
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	// data, ok := response["data"].(map[string]interface{})
	// if !ok {
	// 	t.Error("Failed to assert 'data' field as a map of string-key and interface-value pairs.")
	// }

	// items, ok := data["items"].([]interface{})
	// assert.True(t, ok, "items must be a slice of interfaces")

	// assert.Len(t, items, 2)
}

func TestGetItemValidID(t *testing.T) {
	// handler := NewItemHandler()
	// router := gin.New()
	// router.GET("/items/:itemID", handler.GetItem)

	// req, err := http.NewRequest(http.MethodGet, "/items/1", nil)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

	// var response map[string]interface{}
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))

	// data, ok := response["data"].(map[string]interface{})
	// if !ok {
	// 	t.Error("Failed to assert 'data' field as a map of string-key and interface-value pairs.")
	// }

	// id, ok := data["id"].(float64)
	// assert.True(t, ok, "id must be a float64")
	// assert.Equal(t, 1, int(id))
}

func TestGetItemInvalidID(t *testing.T) {
	// handler := NewItemHandler()
	// router := gin.New()
	// router.GET("/items/:itemID", handler.GetItem)

	// req, err := http.NewRequest(http.MethodGet, "/items/abc", nil)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// var response map[string]interface{}
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	// assert.Equal(t, "invalid item_id", response["message"])
}

func TestGetItemNonExistentID(t *testing.T) {
	// handler := NewItemHandler()
	// router := gin.New()
	// router.GET("/items/:itemID", handler.GetItem)

	// req, err := http.NewRequest(http.MethodGet, "/items/47", nil)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusNotFound, w.Code)

	// var response map[string]interface{}
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	// assert.Equal(t, "item not found", response["message"])
}