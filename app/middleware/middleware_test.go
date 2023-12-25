package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"harryd.com/tools/app/routes"
)

type MockHomeHandler struct{}

func (h *MockHomeHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Mock Home Handler"})
}

type MockItemsHandler struct{}

func (h *MockItemsHandler) GetItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"item_id": "123"})
}

func (h *MockItemsHandler) GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"items": []map[string]interface{}{}})
}

func TestLoggerMiddleware(t *testing.T) {
	var buf bytes.Buffer
	gin.DefaultWriter = &buf
	router := gin.Default()

	homeHandler := &MockHomeHandler{}
	itemsHandler := &MockItemsHandler{}

	routerInitializer := &routes.RouterInitializer{
        MiddlewareInterface: &MiddlewareImpl{},
    }
	routerInitializer.InitializeRoutes(router, homeHandler, itemsHandler)

	req, err := http.NewRequest("GET", "/api/v1/", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())
}
