package middleware

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"*"}
	
	// return cors.New(config)
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		}
	}
}