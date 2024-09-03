package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	validationErrors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		// Add each validation error to the map with the field name as the key
		validationErrors[err.Field()] = fmt.Sprintf("failed on the '%s' tag", err.Tag())
	}

	return validationErrors
}
