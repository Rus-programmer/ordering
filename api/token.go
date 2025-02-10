package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/util"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// createCustomer handles customer creation.
// @Summary Create a new customer
// @Description Register a new customer with username, password, and role
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dto.CreateCustomerRequest true "Customer details"
// @Success 200 {object} dto.CustomerResponse
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
