package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name" binding : "required"`
	Email string `json: "email" binding : "required"`
	Age   int    `json:"age" binding:"required"`
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

	r.POST("/users", func(ctx *gin.Context) {
		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(404, gin.H{
				"err": err.Error()})
			return
		} //db.Create(&user) inserts a row in the database, it doesn’t create a template.
		db.Create(&user)
		ctx.JSON(201, user)

	})
	r.GET("/users", func(ctx *gin.Context) {

		var users []User

		db.Find(&users)
		ctx.JSON(400, gin.H{
			"let see ":   "seeing",
			"users list": users,
		})
	})
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var user User
		result := db.First(&user, id)

		if result.Error != nil {
			ctx.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"user": user,
		})
	})

	r.Run(":9000")

}
