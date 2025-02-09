package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/middleware"
	order "ordering/services/orders"
	"ordering/token"
	"ordering/util"
)

type getOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)

	arg := order.GetOrder{
		ID:      req.ID,
		Payload: payload,
	}

	product, err := server.service.GetOrder(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}
