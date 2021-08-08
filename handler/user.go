package handler

import (
	"fmt"
	"moyu/auth"
	"moyu/helper"
	"moyu/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIResponse("Register account failed", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormatter := user.FormatUser(newUser, token)
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

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userFormatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Successfuly loggedin", "success", http.StatusOK, userFormatter)

	c.JSON(http.StatusOK, response)
}

// TODO: Check available email
func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	// tangkap input dari user
	// dan mapping ke struct CheckEmailInput
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// user input passing ke service
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// response check email available success
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// TODO: Upload Avatar
func (h *userHandler) UploadAvatar(c *gin.Context) {
	// tangkap nput dari user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// dapat currentUser dari middleware
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// simpan avatar di folder "images/"
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// di service kita call repo
	// repo ambil data user yg ID = 1
	// repo update data user simpan lokasi file ke db
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}

		response := helper.APIResponse("Failed to upload avatar image", "error", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", "success", http.StatusOK, data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	curentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(curentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", "success", http.StatusOK, formatter)

	c.JSON(http.StatusOK, response)
}
