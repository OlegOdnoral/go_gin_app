package authhandlers

type signinHandlerModel struct {
	Login    string `json:"login" binding:"required,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=15"`
}
