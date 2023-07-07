package controller

import (
	"gin-restfull-api/data/response"
	"gin-restfull-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UsersRepository
}

func NewUsersController(repository repository.UsersRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
