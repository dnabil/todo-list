package repo

import (
	"todo-list-be/dto"
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

func (r *TodoRepo) IndexByUserWithPage(db *gorm.DB, req *dto.IndexByUserTodoRequest, page *dto.PageMetadata) ([]model.Todo, error){
	var todos []model.Todo
	filter := r.filterTodo(req.UserID, req.IsDone)
	tx := db.Scopes(paginateScope(page.Page, page.Size), filter).Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}

	err := r.Repo.populatePageResult(tx, &todos, page, filter)
	if err != nil {
		return nil , err
	}
	return todos, nil
}

func (r *TodoRepo) filterTodo(userId uint, isDone *bool) func(*gorm.DB) *gorm.DB{
	return func(tx *gorm.DB) *gorm.DB {
		if userId != 0 {
			tx = tx.Where("user_id = ?", userId)
		}

		if isDone != nil{
			tx = tx.Where("is_done = ?", *isDone)
		}

		return tx
	}
}