package repo

import (
	"todo-list-be/model"

	"github.com/sirupsen/logrus"
)

type TodoRepo struct {
	Repo[model.Todo]
	Log *logrus.Logger
}

func NewTodoRepo(log *logrus.Logger) *TodoRepo {
	return &TodoRepo{
		Log: log,
	}
}