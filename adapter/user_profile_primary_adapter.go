package adapter

import (
	"hexagonal/practice/model"
	"hexagonal/practice/port"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserProfilePrimaryAdapter struct {
	service port.UserProfilePrimaryPort
}

// InstanceUserProfilePrimaryAdapter creates a new instance of UserProfilePrimaryAdapter
func InstanceUserProfilePrimaryAdapter(service port.UserProfilePrimaryPort) *UserProfilePrimaryAdapter {
	return &UserProfilePrimaryAdapter{service: service}
}

func (u *UserProfilePrimaryAdapter) CreateUserProfile(c *fiber.Ctx) error {
	var profile model.UserProfile

	// Parse request body
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call service to create user profile
	createdProfile, err := u.service.CreateUserProfile(&profile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User profile created successfully",
		"data":    createdProfile,
	})
}

func (u *UserProfilePrimaryAdapter) GetUserProfileByUserName(c *fiber.Ctx) error {
	// Get user name from URL parameter
	userName := c.Params("userName")
	if userName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User name is required",
		})
	}

	// Call service to get user profile by user name
	profile, err := u.service.GetUserProfileByUserName(userName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User profile retrieved successfully",
		"data":    profile,
	})
}

func (u *UserProfilePrimaryAdapter) UpdateUserProfile(c *fiber.Ctx) error {
	// Get profile ID from URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid profile ID",
		})
	}

	var profile model.UserProfile
	// Parse request body
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call service to update user profile
	updatedProfile, err := u.service.UpdateUserProfile(uint(id), &profile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User profile updated successfully",
		"data":    updatedProfile,
	})
}

func (u *UserProfilePrimaryAdapter) DeleteUserProfile(c *fiber.Ctx) error {
	// Get profile ID from URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid profile ID",
		})
	}

	// Call service to delete user profile
	if err := u.service.DeleteUserProfile(uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User profile deleted successfully",
	})
}

func (u *UserProfilePrimaryAdapter) SearchUserByName(c *fiber.Ctx) error {
	// Get search query from URL parameter
	name := c.Params("name")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name query parameter is required",
		})
	}

	// Call service to search user by name
	profiles, err := u.service.SearchUserByName(name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Search completed successfully",
		"data":    profiles,
	})
}

func (u *UserProfilePrimaryAdapter) GetAllUserProfiles(c *fiber.Ctx) error {
	// Get all user profiles
	profiles, err := u.service.GetAllUserProfiles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user profiles",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User profiles retrieved successfully",
		"data":    profiles,
	})
}
