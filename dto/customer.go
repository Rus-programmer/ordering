package dto

import (
	db "ordering/db/sqlc"
	"time"
)

type CreateCustomerRequest struct {
	Username string      `json:"username" binding:"required,alphanum"`
	Password string      `json:"password" binding:"required,min=6"`
	Role     db.UserRole `json:"role" binding:"required,role"`
}

type CustomerResponse struct {
	ID        int64       `json:"id"`
	Username  string      `json:"username"`
	Role      db.UserRole `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
