package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"reflect"
)

var NotEmptyArrayValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()

	// Проверяем, является ли поле слайсом или массивом
	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		log.Info().Msgf("Field value: %v", field.Len()) // Лог для отладки
		return field.Len() > 0
	}

	return false
}
