package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateStruct[T any](data T) ([]ErrorResponse, error) {
	var errors []ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			// element.Field = FindJsonTagName(err, err.Field())
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, element)
		}
	}
	return errors, err
}
