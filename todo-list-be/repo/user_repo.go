package repo

import (
	"todo-list-be/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *UserRepo) FindByEmail(db *gorm.DB, email string, user *model.User) error{
	return db.Where("email = ?", email).Take(user).Error
}