package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
)

func Validate(request model.StudentCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateUsername(err error) {
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
