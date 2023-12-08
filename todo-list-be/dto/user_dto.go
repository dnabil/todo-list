package dto

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt/v4"
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

type LoginUserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func (r *LoginUserRequest) Validate() error{
	return validation.ValidateStruct(r, 
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 0)),
	)
}

type JwtUserClaims struct{
	jwt.RegisteredClaims
}
func NewJwtUserClaims(id uint, t time.Duration) (JwtUserClaims){
	return JwtUserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: fmt.Sprint(id),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
		},
	}
}

type LoginUserResponse struct{
	Token string `json:"token"`
}