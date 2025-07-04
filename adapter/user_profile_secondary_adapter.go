package adapter

import (
	"fmt"
	"hexagonal/practice/model"
	"hexagonal/practice/port"

	"gorm.io/gorm"
)

type UserProfileSecondaryAdapter struct {
	db *gorm.DB
}

// InstanceUserProfileSecondaryAdapter creates a new instance of UserProfileSecondaryAdapter
func InstanceUserProfileSecondaryAdapter(db *gorm.DB) port.UserProfileSecondaryPort {
	return &UserProfileSecondaryAdapter{db: db}
}

func (u *UserProfileSecondaryAdapter) CreateUserProfile(profile *model.UserProfile) (*model.UserProfile, error) {
	// Check if the user exists before creating profile
	var user model.User
	if err := u.db.First(&user, *profile.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID '%d' not found", *profile.UserID)
		}
		return nil, fmt.Errorf("error finding user: %v", err)
	}

	// Log the user found for debugging
	fmt.Printf("User found: ID=%v, Name=%s, Position=%s\n",
		*user.ID, *user.Name, *user.Position)

	// Create the profile
	if err := u.db.Create(profile).Error; err != nil {
		return nil, fmt.Errorf("error creating profile: %v", err)
	}

	// Log the created profile for debugging
	fmt.Printf("Profile created: ID=%v, UserID=%d\n",
		*profile.ID, *profile.UserID)

	return profile, nil
}

func (u *UserProfileSecondaryAdapter) GetUserProfileByUserName(userName string) (*model.UserProfile, error) {
	var profile model.UserProfile

	// First get the user ID from user name
	var user model.User
	if err := u.db.Where("name = ?", userName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with name '%s' not found", userName)
		}
		return nil, fmt.Errorf("error finding user: %v", err)
	}

	// Then query UserProfile by UserID
	if err := u.db.Where("user_id = ?", *user.ID).First(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func (u *UserProfileSecondaryAdapter) UpdateUserProfile(id uint, profile *model.UserProfile) (*model.UserProfile, error) {
	var existingProfile model.UserProfile
	if err := u.db.First(&existingProfile, id).Error; err != nil {
		return nil, err
	}

	// Update only non-nil fields
	if profile.SkilledLanguage != nil {
		existingProfile.SkilledLanguage = profile.SkilledLanguage
	}
	if profile.Project1 != nil {
		existingProfile.Project1 = profile.Project1
	}
	if profile.Project2 != nil {
		existingProfile.Project2 = profile.Project2
	}
	if profile.Project3 != nil {
		existingProfile.Project3 = profile.Project3
	}

	if err := u.db.Save(&existingProfile).Error; err != nil {
		return nil, err
	}
	return &existingProfile, nil
}

func (u *UserProfileSecondaryAdapter) DeleteUserProfile(id uint) error {
	var profile model.UserProfile
	if err := u.db.First(&profile, id).Error; err != nil {
		return err
	}

	return u.db.Delete(&profile).Error
}

func (u *UserProfileSecondaryAdapter) SearchUserByName(name string) ([]*model.UserProfile, error) {
	var profiles []*model.UserProfile

	// Search user profiles using JOIN with users table
	if err := u.db.Table("user_profiles").
		Select("user_profiles.*").
		Joins("JOIN users ON user_profiles.user_id = users.id").
		Where("users.name LIKE ?", "%"+name+"%").
		Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}

func (u *UserProfileSecondaryAdapter) GetAllUserProfiles() ([]*model.UserProfile, error) {
	var profiles []*model.UserProfile
	if err := u.db.Find(&profiles).Error; err != nil {
		return nil, err
	}

	return profiles, nil
}
