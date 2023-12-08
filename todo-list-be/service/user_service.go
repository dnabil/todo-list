package service

import (
	"context"
	"net/http"
	"os"
	"time"
	"todo-list-be/dto"
	"todo-list-be/helper/errcode"
	"todo-list-be/helper/jwtauth"
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

func (s *UserService) Login(ctx context.Context, req *dto.LoginUserRequest) (string, errcode.ErrCodeI){
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	
	errUnathorized := errcode.New("wrong email/password", http.StatusUnauthorized)

	// retrieve user data (db)
	user := new(model.User)
	err := s.Repo.FindByEmail(tx, req.Email, user)
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return "", errUnathorized
		}

		s.Log.Errorln("failed to search user email:", err)
		return "", errcode.ErrInternalServer
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		s.Log.Warnln("failed to compare password:", err)
		return "", errUnathorized
	}

	// create jwt token
	
	// jwt ttl
	jwtTtl, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		s.Log.Errorln("failed to parse jwt ttl:", err)
		return "", errcode.ErrInternalServer
	}

	// token
	claims := dto.NewJwtUserClaims(user.Username, jwtTtl)
	key := os.Getenv("JWT_KEY")
	if key == "" {
		s.Log.Errorln("jwt key is not set")
		return "", errcode.ErrInternalServer
	}
	token, err := jwtauth.NewToken(claims, []byte(key))
	if err != nil {
		s.Log.Errorln("failed to create jwt token:", err)
		return "", errcode.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.Warnf("Failed commit transaction : %+v\n", err)
		return "", errcode.ErrInternalServer
	}

	return token, nil
}