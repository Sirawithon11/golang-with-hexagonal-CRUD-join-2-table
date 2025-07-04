package core

import (
	"errors"
	"hexagonal/practice/model"
	"hexagonal/practice/port"
)

type UserCore struct {
	secondary port.UserSecondaryPort
}

// NewUserCore creates a new instance of UserCore
func NewUserCore(secondary port.UserSecondaryPort) port.UserPrimaryPort {
	return &UserCore{secondary: secondary}
}

func (u *UserCore) CreateUser(user *model.User) (*model.User, error) {
	// Business logic validation
	if user.Name == nil || *user.Name == "" {
		return nil, errors.New("user name is required")
	}
	if user.Position == nil || *user.Position == "" {
		return nil, errors.New("user position is required")
	}
	ArrayAllUser, err := u.secondary.GetUsers()
	if len(ArrayAllUser) == 5 || err != nil {
		return nil, errors.New("number of users must be less than 5 Or GetUser is problem")

	}

	// Call secondary adapter to create user
	return u.secondary.CreateUser(user)
}

func (u *UserCore) GetUsers() ([]*model.User, error) {
	// Call secondary adapter to get all users
	return u.secondary.GetUsers()
}

func (u *UserCore) GetUserByID(id uint) (*model.User, error) {
	// Business logic validation
	if id == 0 {
		return nil, errors.New("user ID is required")
	}

	// Call secondary adapter to get user by ID
	return u.secondary.GetUserByID(id)
}

func (u *UserCore) UpdateUser(id uint, user *model.User) (*model.User, error) {
	// Business logic validation
	if id == 0 {
		return nil, errors.New("user ID is required")
	}

	// Optional: Check if user exists first
	_, err := u.secondary.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Validate update data if provided
	if user.Name != nil && *user.Name == "" {
		return nil, errors.New("user name cannot be empty")
	}
	if user.Position != nil && *user.Position == "" {
		return nil, errors.New("user position cannot be empty")
	}

	// Call secondary adapter to update user
	return u.secondary.UpdateUser(id, user)
}

func (u *UserCore) DeleteUser(id uint) error {
	// Business logic validation
	if id == 0 {
		return errors.New("user ID is required")
	}

	// Optional: Check if user exists first
	_, err := u.secondary.GetUserByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Call secondary adapter to delete user
	return u.secondary.DeleteUser(id)
}
