package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"publishedYear"`
	Price         int    `json:"price"`
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Book{})

	log.Println("Database connected successfully!")

	r := gin.Default()

	r.POST("/book", posting)
	r.GET("/books", geting)
	r.GET("/book/:id", getbyid)
	r.PUT("/book/:id", putbyid)
	r.PATCH("/book/:id", patching)
	r.DELETE("/book/:id", deleting)

	r.Run(":7990")
}

// POST /book
func posting(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// GET /books
func geting(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		log.Println("DB error:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println("Books fetched from DB:", books)
	c.JSON(200, books)
}

// GET /book/:id
func getbyid(c *gin.Context) {
	var book Book
	id := c.Param("id")

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// PUT /book/:id → full update
func putbyid(c *gin.Context) {
	var book Book
	id := c.Param("id")

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	book.PublishedYear = input.PublishedYear
	book.Price = input.Price

	db.Save(&book)
	c.JSON(http.StatusOK, book)
}

// PATCH /book/:id → partial update
func patching(c *gin.Context) {
	var book Book
	id := c.Param("id")

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&book).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DELETE /book/:id
func deleting(c *gin.Context) {
	var book Book
	id := c.Param("id")

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "deleted successfully",
		"deleted_book": book,
	})
}
