package handler

import (
	"net/http"
	"todo-list-be/dto"
	"todo-list-be/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Log *logrus.Logger
	Service *service.UserService
}

func NewUserHandler(log *logrus.Logger, service *service.UserService) *UserHandler{
	return &UserHandler{
		Log: 		log,
		Service: 	service,
	}
}

func (h *UserHandler) Create(c *gin.Context){
	req := new(dto.CreateUserRequest)
	if err := c.BindJSON(&req); err != nil {
		response(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	if err := req.Validate(); err != nil {
		response(c, http.StatusBadRequest, "validation fail", err)
		return
	}

	user, err := h.Service.Create(c.Request.Context(), req)
	if err != nil {
		response(c, err.Code(), err.Error(), nil)
		return
	}

	response(c, http.StatusCreated, "user created", user)
}

func (h *UserHandler) Login(c *gin.Context){
	req := new(dto.LoginUserRequest)
	if err := c.BindJSON(&req); err != nil {
		h.Log.Warnln("bad request:", err)
		response(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	if err := req.Validate(); err != nil {
		response(c, http.StatusBadRequest, "validation fail", err)
		return
	}

	token, err := h.Service.Login(c.Request.Context(), req)
	if err != nil {
		response(c, err.Code(), err.Error(), nil)
		return
	}

	response(c, http.StatusOK, "login success", dto.LoginUserResponse{Token: token})
}