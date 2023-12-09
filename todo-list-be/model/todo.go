package model

import "time"

type Todo struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	IsDone bool `json:"is_done"`

	User User `json:"user,omitempty"` // belongs to 
}