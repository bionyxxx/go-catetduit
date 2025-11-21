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
			case "eqfield":
				errorMessages[fieldName] = fmt.Sprintf("Must be equal to %s", strings.ToLower(fieldError.Param()))
			case "nefield":
				errorMessages[fieldName] = fmt.Sprintf("Must not be equal to %s", strings.ToLower(fieldError.Param()))
			case "gte":
				errorMessages[fieldName] = fmt.Sprintf("Must be greater than or equal to %s", fieldError.Param())
			case "lte":
				errorMessages[fieldName] = fmt.Sprintf("Must be less than or equal to %s", fieldError.Param())
			case "numeric":
				errorMessages[fieldName] = "Must be a numeric value"
			case "url":
				errorMessages[fieldName] = "Must be a valid URL"
			case "oneof":
				errorMessages[fieldName] = fmt.Sprintf("Must be one of the following values: %s", fieldError.Param())
			// Custom validators
			case "phone":
				// must be 08xx, +62xxx, 62xxx
				errorMessages[fieldName] = "Must be a valid phone number (e.g., 08xx, +62xxx, 62xxx)"
			case "exists":
				errorMessages[fieldName] = "The specified value does not exist"
			case "unique":
				errorMessages[fieldName] = "The specified value must be unique"
			case "old_password":
				errorMessages[fieldName] = "The old password is incorrect"

			default:
				errorMessages[fieldName] = "This field is invalid"
			}
		}
	}

	return errorMessages
}
