package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/models"
)

type HomeHandlerImpl struct{}

func NewHomeHandler() *HomeHandlerImpl {
	return &HomeHandlerImpl{}
}

func (h *HomeHandlerImpl) Home(ctx *gin.Context) {
	response := models.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    gin.H{"motto": "Let's Go!"},
	}
	ctx.JSON(http.StatusOK, response)
}
