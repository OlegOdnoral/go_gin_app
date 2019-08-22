package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandleTest handle get request on /auth/test
func HandleTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Done",
	})
}

//SigninHandler handle post request on /auth/signin
func SigninHandler(c *gin.Context) {
	var user signinHandlerModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	fmt.Printf("%+v\n", user)
	c.JSON(200, gin.H{
		"message": "Done",
	})
}

//SignupHandler handle post request on /auth/signup
func SignupHandler(c *gin.Context) {
	var newUser signupHandlerModel
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	fmt.Printf("%+v\n", newUser)
	c.JSON(200, gin.H{
		"message": "Done",
	})
}

//ForgotPasswordHandler handle post request on /auth/forgot-password
func ForgotPasswordHandler(c *gin.Context) {
	var dropPaswordFoUser forgotPasswordMode
	if err := c.ShouldBindJSON(&dropPaswordFoUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	fmt.Printf("%+v\n", dropPaswordFoUser)
	c.JSON(200, gin.H{
		"message": "Done",
	})
}
