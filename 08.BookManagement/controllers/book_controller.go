package controllers

import (
	"fmt"
	"strconv"

	"08.BookManagement/config"
	"08.BookManagement/models"
	"github.com/gin-gonic/gin"
)

// ======= UTILS ==========
func GetBook(c *gin.Context, book *models.Book) (int, *string) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := "Invalid todo id!"
		return 404, &msg
	}

	if err := config.DB.Find(&book, id).Error; err != nil {
		msg := fmt.Sprintf("Failed to fetch book with id %v: %v", id, err.Error())
		return 500, &msg
	}

	if book.ID == 0 {
		msg := fmt.Sprintf("Book with id %v not found!", id)
		return 404, &msg
	}

	return 200, nil
}

// ============= Handlers
func FetchAllBooks(c *gin.Context) {
	books := make([]models.Book, 0)

	if err := config.DB.Find(&books).Error; err != nil {
		c.JSON(500, gin.H{"msg": fmt.Sprintf("Failed to fetch books: %v", err.Error())})
		return
	}

	c.JSON(200, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid body!"})
		return
	}

	if book.Title == "" || book.Author == "" || book.PublishedYear == 0 {
		c.JSON(400, gin.H{"msg": "Some fields are missing!"})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(500, gin.H{"msg": fmt.Sprintf("Failed to create book: %v", err.Error())})
		return
	}

	c.JSON(201, book)
}

func FetchBook(c *gin.Context) {
	var book models.Book
	statusCode, msg := GetBook(c, &book)

	if msg != nil {
		c.JSON(statusCode, gin.H{"msg": *msg})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	statusCode, msg := GetBook(c, &book)

	if msg != nil {
		c.JSON(statusCode, gin.H{"msg": *msg})
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		c.JSON(500, gin.H{"msg": fmt.Sprintf("Failed to delete book with id %v: %v", book.ID, err.Error())})
		return
	}

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	var dbBook models.Book
	statusCode, msg := GetBook(c, &dbBook)

	if msg != nil {
		c.JSON(statusCode, gin.H{"msg": *msg})
		return
	}

	var book models.Book
	if err := c.ShouldBindBodyWithJSON(&book); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid body!"})
		return
	}

	if err := config.DB.Model(&dbBook).Updates(&book).Error; err != nil {
		c.JSON(500, gin.H{"msg": fmt.Sprintf("Failed to update book with id %v: %v", book.ID, err.Error())})
		return
	}

	c.JSON(200, dbBook)
}
