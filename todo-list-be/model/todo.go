package model

import "time"

type Todo struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	IsDone bool
}