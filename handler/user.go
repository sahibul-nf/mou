package handler

import (
	"mou/helper"
	"mou/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// endpoint register
// endpoint available email
// endpoint login
// endpoint upload
// TODO: endpoint "/users"
type userHandler struct {
	userService user.Service
}

func NewUserHandler(s user.Service) *userHandler {
	return &userHandler{s}
}

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
