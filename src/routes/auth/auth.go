package auth

import (
	"github.com/OlegOdnoral/go_gin_app/src/routes/handlers/authhandlers"
	"github.com/gin-gonic/gin"
)

//RegHandlers register handles for auth module
func RegHandlers(r *gin.Engine) {
	auth := r.Group("/auth")
	regGets(auth)
	regPosts(auth)
}

func regPosts(g *gin.RouterGroup) {
	g.POST("/signin", authhandlers.SigninHandler)
	g.POST("/signup", authhandlers.SignupHandler)
	g.POST("/forgot-password", authhandlers.ForgotPasswordHandler)
}

func regGets(g *gin.RouterGroup) {
	g.GET("/test", authhandlers.HandleTest)
}
