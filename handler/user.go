package handler

import (
	"mou/helper"
	"mou/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(s user.Service) *userHandler {
	return &userHandler{s}
}

// TODO: Register
func (h *userHandler) RegisterUser(c *gin.Context) {

	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	var input user.RegisterUserInput

	// fungsi untuk mapping json input ke struct RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Register account failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// struct diatas kita parsing sebagai parameter
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()

	userFormatter := user.FormatUser(newUser, "abc")
	response := helper.APIResponse("Account has been registered", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

// TODO: Login
func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userFormatter := user.FormatUser(loggedInUser, "abcc")
	response := helper.APIResponse("Successfuly loggedin", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

// TODO: Check available email
// TODO: Upload Avatar
