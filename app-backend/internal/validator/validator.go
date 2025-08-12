package validator

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"willdo/internal/config"
)

// ValidationError wraps validator.FieldError
type ValidationError struct {
	validator.FieldError
}

type ValidationErrors []ValidationError

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validator handles struct validation
type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	return &Validator{validate}
}

func (v *Validator) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

func validateDate(fl validator.FieldLevel) bool {
	_, err := time.Parse(config.DateTime, fl.Field().String())
	return err == nil
}
