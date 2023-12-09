package config

import (
	"todo-list-be/app/http/handler"
	"todo-list-be/app/http/middleware"
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
	userRepo := repo.NewUserRepo(cfg.Log)
	todoRepo := repo.NewTodoRepo(cfg.Log)

	userService := service.NewUserService(cfg.DB, cfg.Log, userRepo)
	todoService := service.NewTodoService(cfg.DB, cfg.Log, todoRepo)

	routeCfg := route.RouteConfig{
		App: cfg.App,

		AuthMiddleware: middleware.AuthMiddleware(userService),

		UserHandler: handler.NewUserHandler(cfg.Log, userService),
		TodoHandler: handler.NewTodoHandler(cfg.Log, todoService),
	}

	// load routes
	routeCfg.Load()
}

func NewGin() *gin.Engine{
	app := gin.Default()

	return app
}