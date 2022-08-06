package controllers

import (
	"github.com/GalloaFranco/gin-first-approach/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ILoginController interface {
	Login(context *gin.Context)
}

type loginController struct {
	loginService services.ILoginService
	jwtService   services.IJWTService
}

type Credentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func NewLoginController(loginSerivice services.ILoginService, jwtService services.IJWTService) ILoginController {
	return &loginController{
		loginService: loginSerivice,
		jwtService:   jwtService,
	}
}

// Login Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @Tags auth
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200
// @Failure 401
// @Router /auth/login [post]
func (loginController *loginController) Login(context *gin.Context) {
	var credentials Credentials
	if err := context.ShouldBind(&credentials); err != nil {
		context.JSON(http.StatusUnauthorized, err)
	}
	isAuthenticated := loginController.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		context.JSON(http.StatusOK, loginController.jwtService.GenerateToken(credentials.Username, true))
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
