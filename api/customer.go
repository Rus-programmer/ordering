package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/dto"
	"ordering/services/auth"
	"ordering/services/customers"
	"ordering/util"
)

// createCustomer handles customer creation.
// @Summary Create a new customer
// @Description Register a new customer with username, password, and role
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dto.CreateCustomerRequest true "Customer details"
// @Success 200 {object} dto.CustomerResponse
// @Router /customers [post]
func (server *Server) createCustomer(ctx *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := customers.CreateCustomer{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	customer, err := server.service.CreateCustomer(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

// login handles user authentication.
// @Summary User login
// @Description Authenticate user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Router /login [post]
func (server *Server) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	userAgent := ctx.Request.UserAgent()
	clientIp := ctx.ClientIP()
	loginReq := auth.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	rsp, err := server.service.Login(ctx, loginReq, clientIp, userAgent)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}
