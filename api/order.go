package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/dto"
	"ordering/middleware"
	order "ordering/services/orders"
	"ordering/token"
	"ordering/util"
)

func (server *Server) getOrder(ctx *gin.Context) {
	var req dto.GetOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := order.GetOrder{
		ID:      req.ID,
		Payload: payload,
	}

	orderResponse, err := server.service.GetOrder(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, orderResponse)
}

func (server *Server) listOrders(ctx *gin.Context) {
	var req dto.ListOrderQueries
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := order.ListOrders{
		Payload: payload,
		QueryParams: order.QueryParams{
			MinPrice: req.MinPrice,
			MaxPrice: req.MaxPrice,
			Status:   req.Status,
		},
	}

	orders, err := server.service.ListOrders(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (server *Server) deleteOrder(ctx *gin.Context) {
	var req dto.DeleteOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := order.DeleteOrderParams{
		ID:      req.ID,
		Payload: payload,
	}

	err := server.service.DeleteOrder(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
