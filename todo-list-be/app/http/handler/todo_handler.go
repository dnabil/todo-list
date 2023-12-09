package handler

import (
	"net/http"
	"todo-list-be/dto"
	"todo-list-be/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	Log     *logrus.Logger
	Service *service.TodoService
}

func NewTodoHandler(log *logrus.Logger, service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		Log:     log,
		Service: service,
	}
}

func (h *TodoHandler) Create(c *gin.Context){
	req := new(dto.CreateTodoRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, "bad request", nil)
		return
	}
	
	if err := req.Validate(); err != nil {
		Response(c, http.StatusBadRequest, "validation fail", err)
		return
	}

	auth, _, err := getAuth(c, h.Log)
	if err != nil {
		Response(c, http.StatusUnauthorized, "anauthorized", nil)
		return
	}

	req.UserID = auth.ID

	todo, err := h.Service.Create(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusCreated, "todo created", todo)
}

func (h *TodoHandler) Update(c *gin.Context){
	req := new(dto.UpdateTodoRequest)
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, "bad request", nil)
		return
	}
	
	if err := req.Validate(); err != nil {
		Response(c, http.StatusBadRequest, "validation fail", err)
		return
	}
	
	auth, _, err := getAuth(c, h.Log)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}
	
	req.UserID = auth.ID
	
	todo, err := h.Service.Update(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}
	
	Response(c, http.StatusOK, "update success", todo)
}