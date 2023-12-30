package service

import "github.com/user_management_system/model"

type UserService interface {
	Create(user model.User)
	Update(user model.User)
	Delete(id int)
	GetById(id int) model.User
	GetAll() []model.User
}
