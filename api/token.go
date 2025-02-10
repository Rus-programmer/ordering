package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/util"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// renewAccessToken handles access token renewal.
// @Summary Renew access token
// @Description Refresh the access token using a valid refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body renewAccessTokenRequest true "Refresh token request"
// @Success 200 {object} dto.RenewAccessTokenResponse
// @Router /renew_access [post]
func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	rsp, err := server.service.RenewAccessToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
