package jwtauth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrWrongSigningMethod = errors.New("wrong signing method for token")
)