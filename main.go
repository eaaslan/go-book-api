package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-fiber/models"
	services "go-fiber/service"
	"go-fiber/storage"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	config := &storage.Config{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		Pass:    os.Getenv("DB_PASS"),
		User:    os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName:  os.Getenv("DB_NAME")}

	db, err := storage.NewConnection(config)

	err = models.MigrateBooks(db)
	if err != nil {
		return
	}
	app := fiber.New()

	r := services.Repository{DB: db}
	r.SetupRoutes(app)
	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
