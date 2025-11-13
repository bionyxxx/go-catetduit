package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {

	errorMessages := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			fieldName := strings.ToLower(fieldError.Field())

			tag := fieldError.Tag()

			switch tag {
			case "required":
				errorMessages[fieldName] = "This field is required"
			case "email":
				errorMessages[fieldName] = "Must be a valid email address"
			case "min":
				errorMessages[fieldName] = fmt.Sprintf("Must be at least %s characters long", fieldError.Param())
			case "max":
				errorMessages[fieldName] = fmt.Sprintf("Must be at most %s characters long", fieldError.Param())
			default:
				errorMessages[fieldName] = "This field is invalid"
			}
		}
	}

	return errorMessages
}
