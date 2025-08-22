package main

import (
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
	db.AutoMigrate(&User{})

	log.Println("Database connected successfully!", db)

	r := gin.Default()

	//Create a user

	r.POST("/users", func(ctx *gin.Context) {
		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error()})
			return
		}
		db.Create(&user)
		ctx.JSON(201, user)

	})
	// get all the users
	r.GET("/users", func(ctx *gin.Context) {
		var user []User
		db.Find(&user)
		ctx.JSON(200, user)
	})
	r.Run(":8080")
}
