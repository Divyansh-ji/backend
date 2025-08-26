package main

import (
	"log"
	"net/http"

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

	r.PUT("/products/:id", func(ctx *gin.Context) {
		var user User
		id := ctx.Param("id")

		//find the existing user

		if err := db.First(&user, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"error": "product not found",
			})
			return
		}

		//dekho we find the user

		var input User

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		user.Name = input.Name
		user.Email = input.Email

		db.Save(&user)
		ctx.JSON(200, user)

	})
	//PATCH /products/:id → Update only fields sent in request (delta

	r.PATCH("/products/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user User
		// this step will get us to the pointer we need
		if err := db.First(&user, id).Error; err != nil {

			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}
		// the only catch in this we will take json in the map in interface because json sent by client will be specific so feil that need to be change will be given only
		var input map[string]interface{}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return

		}
		if err := db.Model(&user).Updates(input).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update products",
			})
			return
		}
		ctx.JSON(http.StatusOK, user)

	})

	r.Run(":9040")

}
