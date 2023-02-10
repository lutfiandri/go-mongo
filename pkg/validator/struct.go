package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type StructErrorFields struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct[T any](data T) ([]StructErrorFields, error) {
	var errorFields []StructErrorFields
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorField StructErrorFields

			errorField.Field = err.Field()
			errorField.Tag = err.Tag()
			errorField.Value = err.Param()

			// find json tag
			reflected := reflect.TypeOf(data)
			if reflected.Kind() == reflect.Pointer {
				reflected = reflected.Elem()
			}

			if field, ok := reflected.FieldByName(err.Field()); ok {
				if jsonTag := field.Tag.Get("json"); jsonTag != "" {
					errorField.Field = field.Tag.Get("json")
				}
			}

			errorFields = append(errorFields, errorField)
		}
	}
	return errorFields, err
}
