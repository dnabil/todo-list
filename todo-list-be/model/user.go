package model

import (
	"time"
)

type User struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` 
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`

	Todos []Todo `json:"todos,omitempty"`// has many
}