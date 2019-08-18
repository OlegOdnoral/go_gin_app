package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Done",
	})
}

func SigninHandler(c *gin.Context) {
	var user signinHandlerModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", user)
	c.JSON(200, gin.H{
		"message": "Done",
	})
}

func SignupHandler(c *gin.Context) {

}

func ForgotPasswordHandler(c *gin.Context) {

}
