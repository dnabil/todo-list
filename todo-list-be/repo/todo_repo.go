package repo

import (
	"todo-list-be/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *TodoRepo) FindByIdAndUserId(db *gorm.DB, todo *model.Todo, id uint, userId uint) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(todo).Error
}