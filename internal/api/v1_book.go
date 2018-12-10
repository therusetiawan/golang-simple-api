package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/therusetiawan/golang-simple-api/internal/models"
	"github.com/therusetiawan/golang-simple-api/pkg/db"
)

func getAllBooks(c *gin.Context) {
	var books []models.Book
	err := db.GetConnection().Model(&books).Select() // get all books

	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
	}

	result := map[string]interface{}{
		"books": books,
	}

	httpOkResponse(c, result)
}

func createBook(c *gin.Context) {
	form := &struct {
		Title     string `form:"title" json:"title"`
		Author    string `form:"author" json:"author"`
		Publisher string `form:"publisher" json:"author"`
		Year      string `form:"year" json:"year"`
	}{}
	c.Bind(form)

	// form validation
	err := validation.Errors{
		"title":     validation.Validate(form.Title, validation.Required),
		"author":    validation.Validate(form.Author, validation.Required),
		"publisher": validation.Validate(form.Publisher, validation.Required),
		"year":      validation.Validate(form.Year, validation.Required),
	}.Filter()

	if err != nil {
		httpValidationErrorResponse(c, err.Error())
		return
	}

	book := models.Book{
		Title:     form.Title,
		Author:    form.Author,
		Publisher: form.Publisher,
		Year:      form.Year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.GetConnection().Insert(&book)
	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
		return
	}

	result := map[string]interface{}{
		"books": book,
	}

	httpOkResponse(c, result)
}

func getBookById(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	// get from database
	err := db.GetConnection().Model(&book).Where("id = ?", id).Select()
	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
		return
	}

	result := map[string]interface{}{
		"books": book,
	}

	httpOkResponse(c, result)
}

func updateBookById(c *gin.Context) {
	var book models.Book

	form := &struct {
		Title     string `form:"title" json:"title"`
		Author    string `form:"author" json:"author"`
		Publisher string `form:"publisher" json:"author"`
		Year      string `form:"year" json:"year"`
	}{}
	id := c.Param("id")
	c.Bind(form)

	err := validation.Errors{
		"id":        validation.Validate(id, validation.Required),
		"title":     validation.Validate(form.Title, validation.Required),
		"author":    validation.Validate(form.Author, validation.Required),
		"publisher": validation.Validate(form.Publisher, validation.Required),
		"year":      validation.Validate(form.Year, validation.Required),
	}.Filter()
	if err != nil {
		httpValidationErrorResponse(c, err.Error())
		return
	}

	err = db.GetConnection().Model(&book).Where("id = ?", id).Select()
	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
		return
	}

	book = models.Book{
		Id:        book.Id,
		Title:     form.Title,
		Author:    form.Author,
		Publisher: form.Publisher,
		Year:      form.Year,
	}

	_, err = db.GetConnection().Model(&book).
		Column("title").
		Column("author").
		Column("publisher").
		Column("year").
		WherePK().Returning("*").Update()

	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
		return
	}

	result := map[string]interface{}{
		"book": book,
	}

	httpOkResponse(c, result)
}

func deleteBookById(c *gin.Context) {
	var book models.Book

	id := c.Param("id")
	err := validation.Errors{
		"id": validation.Validate(id, validation.Required),
	}.Filter()
	if err != nil {
		httpValidationErrorResponse(c, err.Error())
		return
	}

	err = db.GetConnection().Model(&book).Where("id = ?", id).Select()
	if err != nil {
		httpInternalServerErrorResponse(c, err.Error())
		return
	}

	db.GetConnection().Delete(&book)

	result := map[string]interface{}{
		"book": "Delete book successfully",
	}

	httpOkResponse(c, result)
}
