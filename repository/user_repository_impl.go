package repository

import (
	"errors"

	"github.com/user_management_system/helper"
	"github.com/user_management_system/model"
	"gorm.io/gorm"
)

// has a dependency on gorm.DB
type UserRepositoryImpl struct {
	Db *gorm.DB
}

// injecting the dependency and creating instance that implements UserRepository
func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// GetAll implements UserRepository.
func (userRepo *UserRepositoryImpl) GetAll() []model.User {
	var users []model.User
	result := userRepo.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// GetById implements UserRepository.
func (userRepo *UserRepositoryImpl) GetById(id int) (user model.User, err error) {
	result := userRepo.Db.Find(&user, id)
	if result.Error != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Delete implements UserRepository.
func (userRepo *UserRepositoryImpl) Delete(id int) {

	var user model.User

	result := userRepo.Db.Where("id = ?", id).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// Save implements UserRepository.
func (userRepo *UserRepositoryImpl) Save(user model.User) {
	result := userRepo.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (userRepo *UserRepositoryImpl) Update(user model.User) {
	result := userRepo.Db.Model(&user).Updates(user)
	helper.ErrorPanic(result.Error)
}
