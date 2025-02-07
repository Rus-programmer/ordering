package api

import (
	"net/http"
	"ordering/util"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
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
