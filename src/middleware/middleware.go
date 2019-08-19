package middleware

import (
	"fmt"

	"github.com/Tempeny/gin_tes/src/database"
	"github.com/gin-gonic/gin"
)

//JWTMiddleware function for check user auth.
//For this it use Authorization request header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(403)
		}
		fmt.Println(c.GetHeader("Authorization"))
		database.PoolStats()
		c.Next()
	}
}
