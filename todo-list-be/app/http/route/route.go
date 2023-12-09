package route

import (
	"net/http"
	"todo-list-be/app/http/handler"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App				*gin.Engine
	UserHandler    	*handler.UserHandler
	TodoHandler	*handler.TodoHandler

	AuthMiddleware	gin.HandlerFunc

}

func (r *RouteConfig) Load() {
	r.App.GET("/", func(c *gin.Context) {c.String(http.StatusOK, "web app is up and running!")})
	
	api := r.App.Group("/api")
	api.POST("/users/register", r.UserHandler.Create)
	api.POST("/users/login", r.UserHandler.Login)
	api.GET("/users/me", r.AuthMiddleware, r.UserHandler.Me)


	todo := api.Group("/todos")
	todo.Use(r.AuthMiddleware)
	todo.POST("/", r.TodoHandler.Create)
	todo.PUT("/", r.TodoHandler.Update)
	// todo.GET("/todos", r.TodoHandler.Index)
	todo.DELETE("/:todoId", r.TodoHandler.Delete)
}