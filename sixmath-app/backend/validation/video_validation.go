package validation

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
)

func VideoValidate(request model.VideoCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.URLVideo, validation.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}

func VideoUpdateValidate(request model.VideoUpdateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.URLVideo, validation.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}
