package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json: "email"`
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully!", db)

	r := gin.Default()
	r.POST("/users", func(ctx *gin.Context) {

		user := User{Name: "Alice", Email: "divyanstiwary01@gmail.com"}

		result := db.Create(&user)
		if result.Error != nil {
			fmt.Println("Error inserting user:", result.Error)
		} else {
			fmt.Println("New user ID:", user.ID)
		}

	})
	r.GET("/users", func(c *gin.Context) {
		var user []User

		db.Find(&user)
		c.JSON(200, user)

	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user User
		if err := db.First(&user, id); err != nil {
			ctx.JSON(400, gin.H{
				"error": "user not found",
			})
			return
		}
		ctx.JSON(200, user)
	})
	//updating inside the database

	r.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user User

		if err := db.First(&user, id); err != nil {
			ctx.JSON(400, gin.H{
				"error": "user not found",
			})
			return
		}
		//Binding new data from request
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error()})
			return
		}

		//save the updates
		db.Save(&user)
		ctx.JSON(200, user)
	})
	r.Run(":7000")
}
