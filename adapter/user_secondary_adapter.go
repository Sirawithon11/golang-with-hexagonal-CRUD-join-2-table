package adapter

import (
	"hexagonal/practice/model"
	"hexagonal/practice/port"

	"gorm.io/gorm"
)

type UserSecondaryAdapter struct {
	db *gorm.DB
}

// InstanceSecondaryAdapter creates a new instance of UserSecondaryAdapter
func InstanceSecondaryAdapter(db *gorm.DB) port.UserSecondaryPort {
	return &UserSecondaryAdapter{db: db}
}

func (u *UserSecondaryAdapter) CreateUser(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserSecondaryAdapter) GetUsers() ([]*model.User, error) {
	var users []*model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserSecondaryAdapter) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserSecondaryAdapter) UpdateUser(id uint, user *model.User) (*model.User, error) {
	var existingUser model.User
	if err := u.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	// Update only non-nil fields
	if user.Name != nil {
		existingUser.Name = user.Name
	}
	if user.Position != nil {
		existingUser.Position = user.Position
	}

	if err := u.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func (u *UserSecondaryAdapter) DeleteUser(id uint) error {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return err
	}

	return u.db.Delete(&user).Error
}
