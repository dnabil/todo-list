package model

import "time"

type User struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Password  string
}