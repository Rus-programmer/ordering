package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/util"
)

// getMetrics handles retrieving system metrics.
// @Summary Get system metrics
// @Description Retrieve various system metrics
// @Tags metrics
// @Accept json
// @Produce json
// @Success 200 {object} dto.MetricsResponse
// @Router /metrics [get]
func (server *Server) getMetrics(ctx *gin.Context) {
	response, err := server.service.GetMetrics(ctx)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
