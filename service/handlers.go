package services

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/models"
	"gorm.io/gorm"
	"log"
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
	api.Get("/get_book/:id", r.GetBook)
	api.Get("/get_books", r.GetBooks)
	api.Delete("/delete_book/:id", r.DeleteBook)
}

func (r *Repository) CreateBook(ctx *fiber.Ctx) error {
	book := &Book{}

	err := ctx.BodyParser(book)
	if err != nil {
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBook(ctx *fiber.Ctx) error {
	book := &models.Books{}
	id := ctx.Params("id")
	if id == "" {
		log.Fatalln("id can not be empty")
	}
	err := r.DB.Where("id = ?", id).First(book)
	log.Fatalln(err)
	return nil
}

func (r *Repository) GetBooks(ctx *fiber.Ctx) error {

	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (r *Repository) DeleteBook(ctx *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := ctx.Params("id")
	if id == "" {
		log.Fatalln("id could not found")
	}
	err := r.DB.Delete(bookModel, id)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
