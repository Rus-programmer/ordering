package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/util"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	rsp, err := server.service.RenewAccessToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
	}

	ctx.JSON(http.StatusOK, rsp)
}
