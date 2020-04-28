package routes

import (
	"github.com/OlegOdnoral/go_gin_app/src/middleware"
	"github.com/OlegOdnoral/go_gin_app/src/routes/auth"
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
