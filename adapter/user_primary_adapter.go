package adapter

import (
	"hexagonal/practice/model"
	"hexagonal/practice/port"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserPrimaryAdapter struct {
	service port.UserPrimaryPort
}

// InstancePrimaryAdapt creates a new instance of UserPrimaryAdapter
func InstancePrimaryAdapt(service port.UserPrimaryPort) *UserPrimaryAdapter {
	return &UserPrimaryAdapter{service: service}
}

func (u *UserPrimaryAdapter) CreateUser(c *fiber.Ctx) error {
	var user model.User

	// Parse request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call service to create user
	createdUser, err := u.service.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    createdUser,
	})
}

func (u *UserPrimaryAdapter) GetUser(c *fiber.Ctx) error {
	// Get all users
	users, err := u.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

func (u *UserPrimaryAdapter) GetUserByID(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Call service to get user by ID
	user, err := u.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func (u *UserPrimaryAdapter) UpdateUser(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var user model.User
	// Parse request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call service to update user
	updatedUser, err := u.service.UpdateUser(uint(id), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}

func (u *UserPrimaryAdapter) DeleteUser(c *fiber.Ctx) error {
	// Get user ID from URL parameter
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Call service to delete user
	if err := u.service.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
