package handler

import (
	"fmt"
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
		fmt.Printf("%v %T\n", err, err)
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

func (h *TodoHandler) Delete(c *gin.Context){
	req := new(dto.DeleteTodoRequest)
	if err := c.ShouldBindUri(req); err != nil {
		Response(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	auth, _, err := getAuth(c, h.Log)
	if err != nil {
		Response(c, http.StatusUnauthorized, "anauthorized", nil)
		return
	}

	req.UserID = auth.ID

	if err := h.Service.Delete(c.Request.Context(), req); err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusOK, "todo deleted", nil)
}

func (h *TodoHandler) IndexByUser(c *gin.Context){
	req := new(dto.IndexByUserTodoRequest)
	if err := c.ShouldBindQuery(req); err != nil{
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

	todos, err := h.Service.IndexByUser(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusOK, "found", todos)
}

func (h *TodoHandler) UpdateIsDone(c *gin.Context){
	req := new(dto.UpdateIsDoneTodoRequest)
	if err := c.ShouldBindJSON(req); err != nil{
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

	todo, err := h.Service.UpdateIsDone(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusOK, fmt.Sprintf("todo updated to %t", *req.IsDone), todo)
}