package order

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/token"
	"ordering/util"
)

type DeleteOrderParams struct {
	ID      int64
	Payload *token.Payload
}

func (o *order) DeleteOrder(ctx context.Context, req DeleteOrderParams) error {
	arg := db.GetOrderParams{
		CustomerID: req.Payload.CustomerID,
		ID:         req.ID,
	}
	order, err := o.store.GetOrder(ctx, arg)
	if err != nil {
		return err
	}

	if order.CustomerID != req.Payload.CustomerID && req.Payload.Role != string(db.UserRoleAdmin) {
		return util.ErrRecordNotFound
	}

	deleteDrg := db.SoftDeleteOrderParams{
		CustomerID: req.Payload.CustomerID,
		ID:         req.ID,
	}
	_, err = o.store.SoftDeleteOrder(ctx, deleteDrg)
	if err != nil {
		return err
	}

	return nil
}
