package handlers

import "github.com/go-playground/validator"

type Validator struct {
	Validator *validator.Validate
}

func (validator *Validator) Validate(i interface{}) error {
	return validator.Validator.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{
		Validator: validator.New(),
	}
}
