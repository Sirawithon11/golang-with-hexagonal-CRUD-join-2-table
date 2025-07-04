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
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	SecondaryAdapter := adapter.InstanceSecondaryAdapter(db)
	UserCore := core.NewUserCore(SecondaryAdapter)
	PrimaryAdapter := adapter.InstancePrimaryAdapt(UserCore)

	app.Post("/users", PrimaryAdapter.CreateUser)
	app.Get("/users", PrimaryAdapter.GetUser)
	app.Get("/users/:id", PrimaryAdapter.GetUserByID)
	app.Patch("/users/:id", PrimaryAdapter.UpdateUser)
	app.Delete("/users/:id", PrimaryAdapter.DeleteUser)
	// Start the server
	app.Listen(":8000")
}
