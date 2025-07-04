package port

import "hexagonal/practice/model"

type UserSecondaryPort interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(id uint, user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}
