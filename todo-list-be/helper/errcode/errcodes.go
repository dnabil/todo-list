package errcode

import "net/http"

var (
	ErrInternalServer = New("Internal Server Error", http.StatusInternalServerError)
	ErrBadRequest = New("Bad Request", http.StatusBadRequest)
	ErrConflict = New("Conflict", http.StatusConflict)

	// etc (add pls)
)