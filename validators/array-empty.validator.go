package validators

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var NotEmptyArrayValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		return field.Len() > 0
	}

	return false
}
