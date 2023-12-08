package middleware

import (
	"net/http"
	"os"
	"strings"
	"todo-list-be/app/http/handler"
	"todo-list-be/dto"
	"todo-list-be/helper/errcode"
	"todo-list-be/helper/jwtauth"
	"todo-list-be/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func AuthMiddleware(userService *service.UserService) gin.HandlerFunc{
	// todo: kalau mau pake viper, tambahin config viper sebagai parameter

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			handler.Response(c, http.StatusUnauthorized, "unauthorized", nil)
			c.Abort()
			return
		}

		prefix := "Bearer "
		token = token[(strings.Index(token, prefix) + len(prefix) ):]

		jwtKey := os.Getenv("JWT_KEY")
		if jwtKey == "" {
			userService.Log.Errorln("jwt key is not set!")
			handler.ErrResponse(c, http.StatusInternalServerError)
			return
		}
		
		claims := new(dto.JwtUserClaims)
		err := jwtauth.DecodeToken(token, claims, []byte(jwtKey))
		if err != nil {
			handler.Response(c, http.StatusUnauthorized, "unauthorized", nil)
			return
		}

		c.Set("auth", claims)
	}
}

// helper, not a middleware
func GetAuth(c *gin.Context, log *logrus.Logger) (*dto.JwtUserClaims, errcode.ErrCodeI){
	claims, ok := c.Get("auth")
	if !ok {
		log.Warnln("tried to get auth data on unauthorized user!")
		return nil, errcode.ErrUnauthorized
	}

	userClaims, ok := claims.(*dto.JwtUserClaims)
	if !ok {
		log.Warnf("token can't be parsed to user jwt claims, type is: %T\n", claims)
		return nil, errcode.ErrUnauthorized
	}

	return userClaims, nil
}