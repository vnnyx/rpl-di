package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
)

func VideoValidate(request model.VideoCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.PlaylistId, validation.Required),
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Deskripsi, validation.Required),
		validation.Field(&request.URLVideo, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func VideoUpdateValidate(request model.VideoUpdateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.PlaylistId, validation.Required),
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Deskripsi, validation.Required),
		validation.Field(&request.URLVideo, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
