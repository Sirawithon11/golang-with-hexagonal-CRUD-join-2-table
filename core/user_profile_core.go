package core

import (
	"errors"
	"hexagonal/practice/model"
	"hexagonal/practice/port"
	"strings"
)

type UserProfileCore struct {
	secondary port.UserProfileSecondaryPort
}

// NewUserProfileCore creates a new instance of UserProfileCore
func NewUserProfileCore(secondary port.UserProfileSecondaryPort) port.UserProfilePrimaryPort {
	return &UserProfileCore{secondary: secondary}
}

func (u *UserProfileCore) CreateUserProfile(profile *model.UserProfile) (*model.UserProfile, error) {
	// Business logic validation
	if profile.UserID == nil || *profile.UserID == 0 {
		return nil, errors.New("user ID is required")
	}
	if profile.SkilledLanguage == nil || *profile.SkilledLanguage == "" {
		return nil, errors.New("skilled programming language is required")
	}

	// Call secondary adapter to create user profile
	return u.secondary.CreateUserProfile(profile)
}

func (u *UserProfileCore) GetUserProfileByUserName(userName string) (*model.UserProfile, error) {
	// Business logic validation
	if strings.TrimSpace(userName) == "" {
		return nil, errors.New("user name is required")
	}

	// Call secondary adapter to get user profile by user name
	return u.secondary.GetUserProfileByUserName(userName)
}

func (u *UserProfileCore) UpdateUserProfile(id uint, profile *model.UserProfile) (*model.UserProfile, error) {
	// Business logic validation
	if id == 0 {
		return nil, errors.New("profile ID is required")
	}

	// Validate update data if provided
	if profile.SkilledLanguage != nil && *profile.SkilledLanguage == "" {
		return nil, errors.New("skilled programming language cannot be empty")
	}

	// Call secondary adapter to update user profile
	return u.secondary.UpdateUserProfile(id, profile)
}

func (u *UserProfileCore) DeleteUserProfile(id uint) error {
	// Business logic validation
	if id == 0 {
		return errors.New("profile ID is required")
	}

	// Call secondary adapter to delete user profile
	return u.secondary.DeleteUserProfile(id)
}

func (u *UserProfileCore) SearchUserByName(name string) ([]*model.UserProfile, error) {
	// Business logic validation
	if strings.TrimSpace(name) == "" { // ฟังก์ชันนี้จะตัดช่องว่าง (spaces, tabs, newlines) ออกจากต้นและท้ายของ string name
		return nil, errors.New("search name cannot be empty")
	}

	// Call secondary adapter to search user by name
	return u.secondary.SearchUserByName(name)
}

func (u *UserProfileCore) GetAllUserProfiles() ([]*model.UserProfile, error) {
	// Call secondary adapter to get all user profiles
	return u.secondary.GetAllUserProfiles()
}
