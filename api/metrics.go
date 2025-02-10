package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/util"
)

func (server *Server) getMetrics(ctx *gin.Context) {
	response, err := server.service.GetMetrics(ctx)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
