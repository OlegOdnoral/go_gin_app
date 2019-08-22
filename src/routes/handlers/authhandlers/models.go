package authhandlers

type signinHandlerModel struct {
	Login    string `json:"login" binding:"required,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=15"`
}

type signupHandlerModel struct {
	Login     string `json:"login" binding:"required,min=3,max=15"`
	Passwords string `json:"passwords" binding:"required,min=6,max=15"`
	UserName  string `json:"user_name" binding:"omitempty,min=3,max=15"`
}

type forgotPasswordMode struct {
	Login string `json:"login" binding:"required,min=3,max=15"`
}
