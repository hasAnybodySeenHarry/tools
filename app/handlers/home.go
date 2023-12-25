package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "harryd.com/tools/app/models/api"
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
