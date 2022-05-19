package validation

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func CreateExamValidate(request model.CreateExamRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.DurationInMinute, validation.Required),
		validation.Field(&request.Image, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
