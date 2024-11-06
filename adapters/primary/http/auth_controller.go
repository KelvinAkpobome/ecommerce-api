package http

import (
	"ecommerce-api/application"
	"ecommerce-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	UserService *application.UserService
}


func NewAuthController(userService *application.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: err.Error(), Response: nil})
		return
	}
	user, err := ctrl.UserService.RegisterUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: "User registration failed", Response: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "User registered successfully",
		Response: user,
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Success: false, Message: err.Error(), Response: nil})
		return
	}

	user, err := ctrl.UserService.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.Response{Success: false, Message: "Invalid credentials", Response: err.Error()})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.IsAdmin)
	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User login successfully",
		Response: token,
	})
}
