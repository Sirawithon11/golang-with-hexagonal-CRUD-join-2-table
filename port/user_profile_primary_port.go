package port

import "hexagonal/practice/model"

type UserProfilePrimaryPort interface {
	CreateUserProfile(profile *model.UserProfile) (*model.UserProfile, error)
	GetUserProfileByUserName(userName string) (*model.UserProfile, error)
	UpdateUserProfile(id uint, profile *model.UserProfile) (*model.UserProfile, error)
	DeleteUserProfile(id uint) error
	SearchUserByName(name string) ([]*model.UserProfile, error)
	GetAllUserProfiles() ([]*model.UserProfile, error)
}
