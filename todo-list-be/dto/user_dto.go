package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error{
	return validation.ValidateStruct(r, 
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 0)),
		validation.Field(&r.Username, validation.Required, validation.Length(3, 0), is.Alphanumeric),
	)
}