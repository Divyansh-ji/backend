package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // no need to write id it will include that

	Name string `json:"name"`
	Age  int
}

func main() {
	dsn := "host=127.0.0.1 user=divyansh password=Divyansh dbname=gin_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
	//inserting the 3-5 users
	users := []User{
		{Name: "raju", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 20},
		{Name: "Rohan", Age: 20},
		{Name: "Virat", Age: 20},
		{Name: "Dhoni", Age: 20},
		{Name: "Rohit", Age: 20},
	}
	db.Create(&users)
	var user User
	//db.First(&user, "name= ?", "Bob")
	//db.Delete(&user)

	//batch deltete

	db.Where("AGE = ?", 20).Delete(&user)

	//

	//Query the soft delete

	var remaininguser []User
	db.Unscoped().Find(&remaininguser)

	fmt.Println("user after the soft deltel")
	for _, u := range remaininguser {
		fmt.Printf("ID:%d , Name : %s , Age: %d\n", u.ID, u.Name, u.Age)
	}

	log.Println("Database connected successfully!", db)

}
