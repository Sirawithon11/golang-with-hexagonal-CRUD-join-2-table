package main

import (
	"hexagonal/practice/adapter"
	"hexagonal/practice/core"
	"hexagonal/practice/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&model.User{}, &model.UserProfile{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize User components
	SecondaryAdapter := adapter.InstanceSecondaryAdapter(db)
	UserCore := core.NewUserCore(SecondaryAdapter)
	PrimaryAdapter := adapter.InstancePrimaryAdapt(UserCore)

	// Initialize UserProfile components
	UserProfileSecondaryAdapter := adapter.InstanceUserProfileSecondaryAdapter(db)
	UserProfileCore := core.NewUserProfileCore(UserProfileSecondaryAdapter)
	UserProfilePrimaryAdapter := adapter.InstanceUserProfilePrimaryAdapter(UserProfileCore)

	// User routes
	app.Post("/users", PrimaryAdapter.CreateUser)
	app.Get("/users", PrimaryAdapter.GetUser)
	app.Get("/users/:id", PrimaryAdapter.GetUserByID)
	app.Patch("/users/:id", PrimaryAdapter.UpdateUser)
	app.Delete("/users/:id", PrimaryAdapter.DeleteUser)

	// UserProfile routes
	app.Post("/profiles", UserProfilePrimaryAdapter.CreateUserProfile)
	app.Get("/profiles", UserProfilePrimaryAdapter.GetAllUserProfiles)
	app.Get("/profiles/user/:userName", UserProfilePrimaryAdapter.GetUserProfileByUserName)
	app.Patch("/profiles/:id", UserProfilePrimaryAdapter.UpdateUserProfile)
	app.Delete("/profiles/:id", UserProfilePrimaryAdapter.DeleteUserProfile)

	// Search route - Search users by name and get their profiles
	app.Get("/search/:name", UserProfilePrimaryAdapter.SearchUserByName)

	// Start the server
	log.Println("Server starting on :8000")
	app.Listen(":8000")
}
