// main.go
package main

import (
	"log"
	"os"

	"github.com/HuguesRomain/letsbookit_api/models"
	"github.com/HuguesRomain/letsbookit_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := "disable"

	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=" + sslmode

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("database connected successfully")
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Shop{})

	app := fiber.New()

	routes.UserRoutes(app, db)
	routes.ShopRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
