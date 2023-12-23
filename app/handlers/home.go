package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {}

func (h *HomeHandler) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Let's Go!"})
}