package main

import (
	"emailn/internal/dominio/campanha"

	"github.com/go-playground/validator/v10"
)

func main() {
	campanha := campanha.Campanha{}
	validate := validator.New()
	err := validate.Struct(campanha)
	if err == nil {
		println("Nenhum error")
	} else {
		validationErros := err.(validator.ValidationErrors)
		for _, v := range validationErros {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is invalid: " + v.Tag())
			case "min":
				println(v.StructField() + " is required with min: " + v.Param())
			case "max":
				println(v.StructField() + " is required with max: " + v.Param())
			case "email":
				println(v.StructField() + " is invalid: " + v.Param())
			}

		}
	}
}
