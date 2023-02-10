package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorFields struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct[T any](data T) ([]ErrorFields, error) {
	var errors []ErrorFields
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorFields

			// element.Field = FindJsonTagName(err, err.Field())
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, element)
		}
	}
	return errors, err
}
