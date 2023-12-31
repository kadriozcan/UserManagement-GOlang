package service

import (
	"time"

	"github.com/user_management_system/helper"
	"github.com/user_management_system/model"
	"github.com/user_management_system/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

// GetAll implements UserService.
func (userService *UserServiceImpl) GetAll() []model.User {
	return userService.UserRepository.GetAll()
}

// GetById implements UserService.
func (userService *UserServiceImpl) GetById(id int) model.User {
	userData, err := userService.UserRepository.GetById(id)
	helper.ErrorPanic(err)
	return userData
}

// Create implements UserService.
func (userService *UserServiceImpl) Create(user model.User) {
	userModel := model.User{
		Username:         user.Username,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		PhoneNumber:      user.PhoneNumber,
		RegistrationDate: time.Now(),
	}
	userService.UserRepository.Save(userModel)
}

// Delete implements UserService.
func (userService *UserServiceImpl) Delete(id int) {
	userService.UserRepository.Delete(id)
}

// Update implements UserService.
func (userService *UserServiceImpl) Update(user model.User) {
	userData, err := userService.UserRepository.GetById(user.ID)
	helper.ErrorPanic(err)
	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.PhoneNumber = user.PhoneNumber
	userService.UserRepository.Update(userData)
}
