package service

import (
	"context"
	"todo-list-be/dto"
	"todo-list-be/helper/errcode"
	"todo-list-be/model"
	"todo-list-be/repo"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TodoService struct {
	DB   *gorm.DB
	Log  *logrus.Logger
	Repo *repo.TodoRepo
}

func NewTodoService(db *gorm.DB, log *logrus.Logger, repo *repo.TodoRepo) *TodoService {
	return &TodoService{
		DB:   db,
		Log:  log,
		Repo: repo,
	}
}

func (s *TodoService) Create(ctx context.Context, req *dto.CreateTodoRequest) (*dto.TodoResource, errcode.ErrCodeI) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo := &model.Todo{
		Name: req.Name,
		UserID: req.UserID,
		IsDone: false,
	}

	err := s.Repo.Create(tx, todo)
	if err != nil {
		s.Log.WithError(err).Warnln("couldn't create todo")
		return nil, errcode.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Warnln("Failed commit transaction")
		return nil, errcode.ErrInternalServer
	}

	resource := dto.NewTodoResource(todo)
	
	return resource, nil
}