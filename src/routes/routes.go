package routes

import (
	"github.com/Tempeny/gin_tes/src/middleware"
	"github.com/Tempeny/gin_tes/src/routes/auth"
	"github.com/gin-gonic/gin"
)

func addMiddlewares(r *gin.Engine) {
	r.Use(middleware.JWTMiddleware())
}

//RunAndServe gin server on port which user send as param
func RunAndServe(port string) {
	r := gin.Default()
	addMiddlewares(r)
	auth.RegHandlers(r)

	r.Run(port)
}
