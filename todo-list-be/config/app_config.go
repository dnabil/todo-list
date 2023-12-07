package config

import (
	"todo-list-be/app/http/handler"
	"todo-list-be/app/http/route"
	"todo-list-be/repo"
	"todo-list-be/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB *gorm.DB
	App *gin.Engine
	Log *logrus.Logger
}

func Bootstrap(cfg *BootstrapConfig) {
	// DEFINE HANDLERS, USECASE/SERVICE, REPOSITORIES HERE:
	UserRepo := repo.NewUserRepo(cfg.Log)
	// TodoRepo := repo.NewTodoRepo()

	UserService := service.NewUserService(cfg.DB, cfg.Log, UserRepo)
	// TodoService := service.NewTodoService()

	routeCfg := route.RouteConfig{
		App: cfg.App,
		UserHandler: handler.NewUserHandler(cfg.Log, UserService),
		// TodoHandler: handler.NewTodoHandler,
	}

	// load routes
	routeCfg.Load()
}

func NewGin() *gin.Engine{
	app := gin.Default()

	return app
}