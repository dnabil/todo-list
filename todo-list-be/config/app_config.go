package config

import (
	"todo-list-be/app/http/route"

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
	// UserRepo := repo.NewUserRepo()
	// TodoRepo := repo.NewTodoRepo()

	// UserService := service.NewUserService()
	// TodoService := service.NewTodoService()
	
	// UserHandler := 
	// TodoHandler := ()
	

	routeCfg := route.RouteConfig{
		App: cfg.App,
		// UserHandler: handler.NewUserHandler(),
		// TodoHandler: handler.NewTodoHandler,
	}

	// load routes
	routeCfg.Load()
}

func NewGin() *gin.Engine{
	app := gin.Default()

	return app
}