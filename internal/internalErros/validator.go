package internalerros

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	vadidate := validator.New()
	err := vadidate.Struct(obj)

	if err == nil {
		return nil
	}

	validationErros := err.(validator.ValidationErrors)
	validationError := validationErros[0]

	field := strings.ToLower(validationError.StructField())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " is required with max " + validationError.Param())
	case "min":
		return errors.New(field + " is required with min " + validationError.Param())
	case "email":
		return errors.New(field + " is invalid " + validationError.Param())
	}
	return nil
}
