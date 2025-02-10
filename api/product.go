package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ordering/services/products"
	"ordering/util"
)

type listProductRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// listProducts handles fetching a list of products.
// @Summary List products
// @Description Retrieve a list of products with pagination
// @Tags products
// @Accept json
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page size"
// @Success 200 {array} dto.ProductResponse
// @Router /products [get]
func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := products.ListProductRequest{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listProduct, err := server.service.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, listProduct)
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// getProduct handles fetching a product by its ID.
// @Summary Get a product by ID
// @Description Retrieve a product's details by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Router /products/{id} [get]
func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	product, err := server.service.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type updateProductRequestParam struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateProductRequestBody struct {
	Name     string `form:"name"`
	Price    int64  `form:"price" binding:"omitempty,min=1"`
	Quantity int64  `form:"quantity" binding:"omitempty,min=1"`
}

// updateProduct handles updating an existing product.
// @Summary Update a product
// @Description Update an existing product's details
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body updateProductRequestBody true "Updated product details"
// @Success 200 {object} dto.ProductResponse
// @Router /products/{id} [put]
func (server *Server) updateProduct(ctx *gin.Context) {
	var reqParam updateProductRequestParam
	if err := ctx.ShouldBindUri(&reqParam); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	var bodyReq updateProductRequestBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	product, err := server.service.UpdateProduct(ctx, reqParam.ID, products.UpdateProduct{
		Name:     bodyReq.Name,
		Price:    bodyReq.Price,
		Quantity: bodyReq.Quantity,
	})
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type deleteProductRequestParam struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// deleteProduct handles deleting a product by its ID.
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 "Product deleted successfully"
// @Router /products/{id} [delete]
func (server *Server) deleteProduct(ctx *gin.Context) {
	var reqParam deleteProductRequestParam
	if err := ctx.ShouldBindUri(&reqParam); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err := server.service.DeleteProduct(ctx, reqParam.ID)
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type createProductRequestBody struct {
	Name     string `form:"name" binding:"required"`
	Price    int64  `form:"price" binding:"required,min=1"`
	Quantity int64  `form:"quantity" binding:"required,min=1"`
}

// createProduct handles creating a new product.
// @Summary Create a product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body createProductRequestBody true "Product details"
// @Success 200 {object} dto.ProductResponse
// @Router /products [post]
func (server *Server) createProduct(ctx *gin.Context) {
	var bodyReq createProductRequestBody
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	product, err := server.service.CreateProduct(ctx, products.CreateProduct{
		Name:     bodyReq.Name,
		Price:    bodyReq.Price,
		Quantity: bodyReq.Quantity,
	})
	if err != nil {
		ctx.JSON(util.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}
