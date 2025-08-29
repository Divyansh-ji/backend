package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Price int    `json:"price"`
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&User{})

	// Insert sample data
	usersData := []User{
		{Name: "Divyansh", Email: "divyanshtiwaryi@01", Price: 10},
		{Name: "Divyansh", Email: "divyanshtiwryi@01", Price: 21},
		{Name: "Divyash", Email: "divyanshtiwaryi@01", Price: 310},
		{Name: "Divansh", Email: "divyanshtiaryi@01", Price: 410},
		{Name: "Divyanh", Email: "divyanshtiwryi@01", Price: 390},
		{Name: "ivyansh", Email: "divyanstaryi@01", Price: 400},
		{Name: "Divansh", Email: "divyanhtiwaryi@01", Price: 490},
	}
	db.Create(&usersData)

	// Fetch users where price > 100 and price < 500
	//	var users []User
	//result := db.Where("price > ? AND price < ?", 100, 500).Find(&users)
	//if result.Error != nil {
	//log.Fatal(result.Error)
	//}
	var userss []User

	// Print users
	// Pagination logic
	page := 2     // Example: Page 2
	pageSize := 3 // Example: 3 records per page
	offset := (page - 1) * pageSize

	paginated := db.Limit(pageSize).Offset(offset).Find(&userss)
	if paginated.Error != nil {
		log.Fatal(paginated.Error)
	}
	for _, u := range userss {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Price: %d\n", u.ID, u.Name, u.Email, u.Price)
	}

	log.Println("Database connected successfully!", db)
}
