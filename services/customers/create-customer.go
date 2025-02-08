package customers

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
)

type CreateCustomer struct {
	Username string
	Password string
	Role     db.UserRole
}

func (c *customer) CreateCustomer(ctx context.Context, req CreateCustomer) (dto.CustomerResponse, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	arg := db.CreateCustomerParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Role:           req.Role,
	}

	customer, err := c.store.CreateCustomer(ctx, arg)
	if err != nil {
		return dto.CustomerResponse{}, err
	}

	return NewCustomerResponse(customer), nil
}

func NewCustomerResponse(customer db.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:        customer.ID,
		Username:  customer.Username,
		Role:      customer.Role,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}
