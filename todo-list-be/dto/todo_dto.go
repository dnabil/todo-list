package dto

import (
	"reflect"
	"time"
	"todo-list-be/model"

	validation "github.com/go-ozzo/ozzo-validation"
)


type TodoResource struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	IsDone bool `json:"is_done"`

	User *model.User `json:"user,omitempty"`
}
func NewTodoResource(e *model.Todo) *TodoResource{
	resource := new(TodoResource)
	resource.ID = e.ID
	resource.UserID = e.UserID
	resource.CreatedAt = e.CreatedAt
	resource.UpdatedAt = e.UpdatedAt
	resource.Name = e.Name
	resource.IsDone = e.IsDone
	isUserEmpty := reflect.DeepEqual(e.User, model.User{})
	if !isUserEmpty{
		resource.User = &e.User
	}
	return resource
}

type CreateTodoRequest struct {
	Name string `json:"name"`
	UserID uint `json:"-"`
}

func (r *CreateTodoRequest) Validate() error{
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, validation.Length(2, 255)),
	)
}

type UpdateTodoRequest struct {
	ID uint `json:"id"`
	UserID uint `json:"-"`

	Name string `json:"name"`
	IsDone *bool `json:"is_done"`
}
func (r *UpdateTodoRequest) Validate() error{
	return validation.ValidateStruct(r,
		validation.Field(&r.ID, validation.Required),
		validation.Field(&r.Name, validation.Required, validation.Length(2, 255)),
		validation.Field(&r.IsDone, validation.Required), // https://github.com/go-ozzo/ozzo-validation/issues/79 jadi pake *
	)
}