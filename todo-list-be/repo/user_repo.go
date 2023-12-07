package repo

import (
	"todo-list-be/model"

	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	Repo[model.User]
	Log *logrus.Logger
}

func NewUserRepo(log *logrus.Logger) *UserRepo{
	return &UserRepo{
		Log: log,
	}
}