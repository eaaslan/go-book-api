package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
}

func (r *Repository) CreateBook(ctx *fiber.Ctx) error {
	return nil
}
