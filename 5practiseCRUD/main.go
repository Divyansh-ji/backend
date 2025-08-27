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
	ID            uint   `gorm:"primaryKey"`
	Title         string `json:"title"`
	Author        string `json: "author"`
	PublishedYear int    `json :"publishedyear"`
	Price         int    `json:"year`
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Book{})

	log.Println("Database connected successfully!", db)

	r := gin.Default()

	r.POST("/book", posting)
	r.GET("/books", geting)
	r.GET("book/:id", getbyid)
	r.PUT("/book/:id", putbyid)
	r.PATCH("/book/:id", patching)
	r.DELETE("/books/:id", delteing)
	r.Run(":7000")

}
func posting(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, err.Error())
		return
	}
	if err := db.Create(&book).Error; err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, book)

}
func geting(c *gin.Context) {
	var books []Book

	if err := db.Find(&books).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, books)
}
func getbyid(c *gin.Context) {
	var books Book
	id := c.Param("id")

	if err := db.First(&books, id).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, books)

}
func putbyid(c *gin.Context) {
	var book Book

	idstr := c.Param("id")

	if err := db.First(&book, idstr).Error; err != nil {
		c.JSON(404, err.Error())
		return

	}
	var input Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(404, err.Error())
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	book.PublishedYear = input.PublishedYear
	book.Price = input.Price
	db.Save(&book)
	c.JSON(200, book)

}
func patching(c *gin.Context) {
	// only specified feild given by the USer will be modifed
	id := c.Param("id")
	var book Book

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return

	}
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Model(&book).Updates(&input).Error; err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, book)

}
func delteing(c *gin.Context) {
	id := c.Param("id")

	var book Book

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(400, err.Error())
		return

	}
	if err := db.Delete(&book, id).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"status":          "deleted successfully",
		"after delteing ": book,
	})

}
