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

// getOrder handles fetching an order by its ID.
// @Summary Get an order by ID
// @Description Retrieve an order's details by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.OrderResponse
// @Security BearerAuth
// @Router /orders/{id} [get]
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

// listOrders handles fetching a list of orders based on query parameters.
// @Summary List orders
// @Description Retrieve a list of orders with optional filters
// @Tags orders
// @Accept json
// @Produce json
// @Param status query string false "Order status"
// @Param min_price query int false "Minimum price"
// @Param max_price query int false "Maximum price"
// @Success 200 {array} dto.OrderResponse
// @Security BearerAuth
// @Router /orders [get]
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

// deleteOrder handles deleting an order by its ID.
// @Summary Delete an order
// @Description Delete an order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 "Order deleted successfully"
// @Security BearerAuth
// @Router /orders/{id} [delete]
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

// createOrder handles creating a new order.
// @Summary Create an order
// @Description Create a new order with the provided products
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequestBody true "Order details"
// @Success 200 {object} dto.OrderResponse
// @Security BearerAuth
// @Router /orders [post]
func (server *Server) createOrder(ctx *gin.Context) {
	var bodyReq dto.CreateOrderRequestBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	products := make([]order.CreateOrderItem, len(bodyReq.Products))
	for i, item := range bodyReq.Products {
		products[i] = order.CreateOrderItem{
			ProductID:     item.ProductID,
			OrderedAmount: item.OrderedAmount,
		}
	}

	newOrder, err := server.service.CreateOrder(ctx, order.CreateOrderParams{
		Payload:  payload,
		Products: products,
	})
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, newOrder)
}

// updateOrder handles updating an existing order.
// @Summary Update an order
// @Description Update an existing order's details
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body dto.UpdateOrderRequestBody true "Updated order details"
// @Success 200 {object} dto.OrderResponse
// @Security BearerAuth
// @Router /orders/{id} [put]
func (server *Server) updateOrder(ctx *gin.Context) {
	var req dto.UpdateOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	var bodyReq dto.UpdateOrderRequestBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	products := make([]order.UpdateOrderItem, len(bodyReq.Products))
	for i, item := range bodyReq.Products {
		products[i] = order.UpdateOrderItem{
			ProductID:     item.ProductID,
			OrderedAmount: item.OrderedAmount,
		}
	}

	updatedOrder, err := server.service.UpdateOrder(ctx, order.UpdateOrderParams{
		OrderID:  req.ID,
		Status:   bodyReq.Status,
		Payload:  payload,
		Products: products,
	})
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedOrder)
}
