package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Userl struct {
	gorm.Model
	Name      string `json:"name"`
	CompanyID uint   //this will be the foreign key
	Company   Company
}
type Company struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	db.AutoMigrate(&Userl{}, &Company{})

	//lets fill some fake data
	company := []Company{
		{Name: "Google"},
		{Name: "Amazon"},
		{Name: "Meta"},
	}

	users := []Userl{
		{Name: "Alice", CompanyID: 1},
		{Name: "Bob", CompanyID: 2},
		{Name: "Charlie", CompanyID: 3},
	}
	db.Create(&company)
	db.Create(&users)

	var information []Userl

	if err := db.Preload("Company", "Name = ?", "Google").Find(&information, "ID = ?", "1").Error; err != nil {
		log.Fatal(err)
	}

	for _, u := range information {
		fmt.Printf("User: %-7s | Company: %s\n", u.Name, u.Company.Name)
	}

}
