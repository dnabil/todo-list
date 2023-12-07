package service

import (
	"context"
	"net/http"
	"todo-list-be/dto"
	"todo-list-be/helper/errcode"
	"todo-list-be/model"
	"todo-list-be/repo"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
	Log *logrus.Logger
	Repo *repo.UserRepo
}

func NewUserService(db *gorm.DB, log *logrus.Logger, repo *repo.UserRepo) *UserService{
	return &UserService{
		DB:     db,
		Log:    log,
		Repo:	repo,
	}
}


func (s *UserService) Create(ctx context.Context, req *dto.CreateUserRequest) (*model.User, errcode.ErrCodeI) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Warnln("Failed to generate bcrypt hash: ", err)
		return nil, errcode.ErrInternalServer
	}
	
	user := &model.User{
		Username: req.Username,
		Email: req.Email,
		Password: string(password),
	}

	if err := s.Repo.Create(tx, user); err != nil{
		if err == gorm.ErrDuplicatedKey{
			return nil, errcode.New("email/username already exists", http.StatusConflict)
		}
		s.Log.Warnln("Failed create user in db:", err)
		return nil, errcode.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v\n", err)
		return nil, errcode.ErrInternalServer
	}

	return user, nil
}