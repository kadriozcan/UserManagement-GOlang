package repository

import "github.com/user_management_system/model"

type UserRepository interface {
	GetAll() []model.User
	GetById(id int) (user model.User, err error)
	Save(user model.User)
	Update(user model.User)
	Delete(id int)
}
