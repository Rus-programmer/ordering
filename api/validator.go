package api

import (
	"github.com/go-playground/validator/v10"
	db "ordering/db/sqlc"
)

var validRole validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if role, ok := fieldLevel.Field().Interface().(db.UserRole); ok {
		switch role {
		case db.UserRoleAdmin, db.UserRoleUser:
			return true
		}
		return false
	}
	return false
}
