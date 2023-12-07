package handler

import (
	"todo-list-be/helper"

	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, code int, message string, data any) {
	var status helper.Status
	
	switch code/100 {
	case 2:
		status = helper.Success
	case 4:
		status = helper.Fail
	default:
		errResponse(c, code)
		return
	}

	c.JSON(code, helper.ApiResponse{
		Status: status,
		Message: message,
		Data: data,
	})
}

func errResponse(c *gin.Context, code int){
	c.AbortWithStatusJSON(code, helper.ApiResponse{
		Status: helper.Error,
		Message: "Internal Server Error",
	})
}