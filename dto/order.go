package dto

import (
	db "ordering/db/sqlc"
	"time"
)

type ListOrderQueries struct {
	Status   db.OrderStatus `form:"status" binding:"omitempty"`
	MinPrice int64          `form:"min_price" binding:"omitempty,number,min=1"`
	MaxPrice int64          `form:"max_price" binding:"omitempty,number,min=1"`
}

type GetOrderRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type DeleteOrderRequest GetOrderRequest

type OrderProductResponse struct {
	Product       ProductResponse `json:"product"`
	OrderedAmount int64           `json:"ordered_amount"`
}

type OrderResponse struct {
	ID         int64                  `json:"id"`
	CustomerID int64                  `json:"customer_id"`
	Status     db.OrderStatus         `json:"status"`
	TotalPrice int64                  `json:"total_price"`
	IsDeleted  bool                   `json:"id_deleted"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	Products   []OrderProductResponse `json:"products"`
}
