package helper

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		error := strings.Split(e.Error(), "Error:")[1]
		errors = append(errors, error)
	}

	return errors
}
