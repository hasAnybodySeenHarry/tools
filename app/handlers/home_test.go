package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	handler := &HomeHandler{}
	router := gin.New()
	router.GET("/", handler.Home)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
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

	motto, ok := data["motto"].(string)
	if !ok {
		t.Error("Failed to assert 'motto' field as string")
		return
	}

	assert.Equal(t, "Let's Go!", motto)
}
