package route

import (
	"net/http"
	"todo-list-be/app/http/handler"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App				*gin.Engine
	UserHandler    	*handler.UserHandler
	// TodoHandler	*handler.TodoHandler

}

func (r *RouteConfig) Load() {
	r.App.GET("/", func(c *gin.Context) {c.String(http.StatusOK, "web app is up and running!")})
	
	api := r.App.Group("/api")
	api.POST("/users/register", r.UserHandler.Create)
	// api.POST("/users/login", r.UserHandler.Login)
	// api.GET("/users/{id}", r.UserHandler.Show)

	// api.POST("/todos", r.TodoHandler.Create)
	// api.GET("/todos", r.TodoHandler.Index) // ada yang completed ada yang ngga (param)
	// api.DELETE("/todos", r.TodoHandler.DELETE)
}