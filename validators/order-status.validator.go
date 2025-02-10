package validators

import (
	"github.com/go-playground/validator/v10"
	db "ordering/db/sqlc"
)

var ValidOrderStatus validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if status, ok := fieldLevel.Field().Interface().(db.OrderStatus); ok {
		switch status {
		case db.OrderStatusPending, db.OrderStatusConfirmed, db.OrderStatusCancelled:
			return true
		}
		return false
	}
	return false
}
