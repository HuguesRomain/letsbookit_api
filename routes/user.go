package routes

import (
	"github.com/HuguesRomain/letsbookit_api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
}
