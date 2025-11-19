package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

var validate *validator.Validate
var db *sqlx.DB

func SetDB(database *sqlx.DB) {
	db = database
}

func NewCustomValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		RegisterCustomValidations()
	}
	return validate
}

func RegisterCustomValidations() {
	err := validate.RegisterValidation("phone", phoneNumberValidator)
	if err != nil {
		fmt.Println("Error registering custom validation:", err)
	}
	err = validate.RegisterValidation("exists", existsValidator())
	if err != nil {
		fmt.Println("Error registering custom validation:", err)
	}
	err = validate.RegisterValidation("unique", uniqueValidator())
	if err != nil {
		fmt.Println("Error registering custom validation:", err)
	}
}

func phoneNumberValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// Format: 08xx atau +62xxx atau 62xxx
	matched, _ := regexp.MatchString(`^(\+62|62|0)8[0-9]{8,11}$`, phone)
	return matched
}

func existsValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		if db == nil {
			return true
		}

		// Parse parameter: "table.column"
		// Example: validate:"exists=users.email"
		param := fl.Param()
		parts := strings.Split(param, ".")

		if len(parts) != 2 {
			fmt.Printf("Invalid exists parameter: %s\n", param)
			return false
		}

		table := parts[0]
		column := parts[1]
		value := fl.Field().Interface()

		// Whitelist
		allowedTables := map[string]bool{"users": true, "products": true}
		allowedColumns := map[string]bool{"email": true, "id": true}

		if !allowedTables[table] || !allowedColumns[column] {
			return false
		}

		var exists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", table, column)
		err := db.Get(&exists, query, value)

		if err != nil {
			fmt.Printf("Error checking existence: %v\n", err)
			return true
		}

		return exists
	}
}

func uniqueValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		if db == nil {
			return true
		}

		param := fl.Param()
		parts := strings.Split(param, ".")

		if len(parts) != 2 {
			return false
		}

		table := parts[0]
		column := parts[1]
		value := fl.Field().Interface()

		allowedTables := map[string]bool{"users": true}
		allowedColumns := map[string]bool{"phone": true, "email": true, "id": true}

		if !allowedTables[table] || !allowedColumns[column] {
			return false
		}

		var exists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", table, column)
		err := db.Get(&exists, query, value)

		if err != nil {
			fmt.Printf("Error checking uniqueness: %v\n", err)
			return true
		}

		return !exists // Valid if NOT exists
	}
}
