package services

import (
	"covid_admission_api/entities"

	"github.com/go-playground/validator"
)

type CustomValidator interface {
	Validate(i interface{}) error
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return entities.ErrorInvalidForm
	}
	return nil
}

func NewValidateService() CustomValidator {
	return &customValidator{validator: validator.New()}
}
