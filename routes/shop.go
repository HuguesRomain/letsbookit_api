package routes

import (
	"github.com/HuguesRomain/letsbookit_api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func ShopRoutes(app *fiber.App, db *gorm.DB) {
	shopHandler := handlers.NewShopHandler(db)

	app.Post("/shops", shopHandler.Create)
	app.Get("/shops/:id", shopHandler.Get)
	app.Get("/shops", shopHandler.GetAll)
}
