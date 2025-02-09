package dto

import (
	db "ordering/db/sqlc"
	"time"
)

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
