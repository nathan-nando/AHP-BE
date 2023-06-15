package validator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func (c CustomValidator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

func New() *validator.Validate {
	return validator.New()
}
