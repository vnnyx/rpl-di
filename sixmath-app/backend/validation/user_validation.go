package validation

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
)

func CreateStudentValidation(request model.StudentCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Handphone, validation.Required),
		validation.Field(&request.Avatar, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.PasswordConfirmation, validation.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}

func CreateTeacherValidation(request model.TeacherCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Handphone, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.PasswordConfirmation, validation.Required),
		validation.Field(&request.Certificate, validation.Required),
		validation.Field(&request.Avatar, validation.Required),
		validation.Field(&request.Age, validation.Required),
		validation.Field(&request.Description, validation.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}

func CreateParentValidation(request model.ParentCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Handphone, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.PasswordConfirmation, validation.Required),
		validation.Field(&request.Avatar, validation.Required),
		validation.Field(&request.StudentUsername, validation.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}
