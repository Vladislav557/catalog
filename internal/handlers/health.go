package handlers

import (
	"github.com/Vladislav557/catalog/internal/models/http/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func (handler *HealthHandler) Health(ctx *gin.Context) {
	response := responses.HealthResponse{Success: true}
	ctx.JSON(http.StatusOK, response)
}
