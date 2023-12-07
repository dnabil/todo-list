package model

import "time"

type Todo struct {
	ID uint
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	IsDone bool

	User User // belongs to
}