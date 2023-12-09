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
		Response(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	if err := req.Validate(); err != nil {
		Response(c, http.StatusBadRequest, "validation fail", err)
		return
	}

	user, err := h.Service.Create(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusCreated, "user created", user)
}

func (h *UserHandler) Login(c *gin.Context){
	req := new(dto.LoginUserRequest)
	if err := c.BindJSON(&req); err != nil {
		h.Log.Warnln("bad request:", err)
		Response(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	if err := req.Validate(); err != nil {
		Response(c, http.StatusBadRequest, "validation fail", err)
		return
	}

	token, err := h.Service.Login(c.Request.Context(), req)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusOK, "login success", dto.LoginUserResponse{Token: token})
}

func (h *UserHandler) Me(c *gin.Context){
	auth, _, err := getAuth(c, h.Log)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	user, err := h.Service.FindById(c.Request.Context(), auth.ID)
	if err != nil {
		Response(c, err.Code(), err.Error(), nil)
		return
	}

	Response(c, http.StatusOK, "success", user)
}