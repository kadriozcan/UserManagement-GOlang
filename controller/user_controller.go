package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user_management_system/helper"
	"github.com/user_management_system/model"
	"github.com/user_management_system/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(servie service.UserService) *UserController {
	return &UserController{
		userService: servie,
	}
}

// GetAll Controlller
func (controller *UserController) GetAll(ctx *gin.Context) {
	users := controller.userService.GetAll()
	ctx.JSON(http.StatusOK, users)
}

// GetById Controlller
func (controller *UserController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorPanic(err)
	user := controller.userService.GetById(id)
	ctx.JSON(http.StatusOK, user)

}

// Create Controlller
func (controller *UserController) Create(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	helper.ErrorPanic(err)

	controller.userService.Create(user)
	ctx.JSON(http.StatusOK, gin.H{"status": "User created"})
}

// Update Controlller
func (controller *UserController) Update(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	helper.ErrorPanic(err)

	controller.userService.Update(user)
	ctx.JSON(http.StatusOK, gin.H{"status": "User updated"})
}

// Delete Controlller
func (controller *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	helper.ErrorPanic(err)

	controller.userService.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{"status": "User deleted"})
}
