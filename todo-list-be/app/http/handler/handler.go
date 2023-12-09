package handler

import (
	"strconv"
	"todo-list-be/dto"
	"todo-list-be/helper"
	"todo-list-be/helper/errcode"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Response(c *gin.Context, code int, message string, data any) {
	var status helper.Status
	
	switch code/100 {
	case 2:
		status = helper.Success
	case 4:
		status = helper.Fail
	default:
		ErrResponse(c, code)
		return
	}

	c.JSON(code, helper.ApiResponse{
		Status: status,
		Message: message,
		Data: data,
	})
}

func ErrResponse(c *gin.Context, code int){
	c.AbortWithStatusJSON(code, helper.ApiResponse{
		Status: helper.Error,
		Message: "Internal Server Error",
	})
}

// helper, not a middleware
func getAuth(c *gin.Context, log *logrus.Logger) (*dto.Auth, *dto.JwtUserClaims, errcode.ErrCodeI){
	claims, ok := c.Get("auth")
	if !ok {
		log.Warnln("tried to get auth data on unauthorized user!")
		return nil, nil, errcode.ErrUnauthorized
	}

	userClaims, ok := claims.(*dto.JwtUserClaims)
	if !ok {
		log.Warnf("token can't be parsed to user jwt claims, type is: %T\n", claims)
		return nil, nil, errcode.ErrUnauthorized
	}

	userID, err := strconv.Atoi(userClaims.Subject)
	if err != nil {
		log.WithError(err).Warnln("can't parse auth user id from jwt claims")
		return nil, userClaims, errcode.ErrUnauthorized
	}

	auth := &dto.Auth{
		ID: uint(userID),
	}

	return auth, userClaims, nil
}