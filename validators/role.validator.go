package validators

import (
	"github.com/go-playground/validator/v10"
	db "ordering/db/sqlc"
)

var ValidRole validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if role, ok := fieldLevel.Field().Interface().(db.UserRole); ok {
		switch role {
		case db.UserRoleAdmin, db.UserRoleUser:
			return true
		}
		return false
	}
	return false
}
