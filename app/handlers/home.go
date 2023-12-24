package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/models"
)

type HomeHandler struct {}

func (h *HomeHandler) Home(ctx *gin.Context) {
	response := models.APIResponse{
		Status: http.StatusOK,
		Message: "success",
		Data: gin.H{"motto": "Let's Go!"},
	}
	ctx.JSON(http.StatusOK, response)
}