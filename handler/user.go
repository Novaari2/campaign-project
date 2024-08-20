package handler

import (
	"campaign-project/helper"
	"campaign-project/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.ApiResponse("Account Creation Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil {
		response := helper.ApiResponse("Account Creation Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "token")
	response := helper.ApiResponse("Account Successfully Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}