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

func (s *TodoService) Update(ctx context.Context, req *dto.UpdateTodoRequest) (*dto.TodoResource, errcode.ErrCodeI) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	//
	// find by id and userid
	todo := new(model.Todo)
	if err := s.Repo.FindByIdAndUserId(tx, todo, req.ID, req.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcode.ErrNotFound
		}

		s.Log.WithError(err).Errorln("failed to find todo by id and user id")
		return nil, errcode.ErrInternalServer
	}

	// update
	todo.Name = req.Name
	todo.IsDone = *req.IsDone
	if err := s.Repo.Update(tx, todo); err != nil {
		s.Log.WithError(err).Errorln("failed to update(s) todo")
		return nil, errcode.ErrInternalServer
	}
	//

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Warnln("Failed commit transaction")
		return nil, errcode.ErrInternalServer
	}

	resource := dto.NewTodoResource(todo)
	return resource, nil
}