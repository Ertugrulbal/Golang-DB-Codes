package main

// go get -u github.com/jinzhu/gorm
// go get -u github.com/lib/pq

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {

	// preperation
	db, err := gorm.Open("postgres", "user=postgres password=1234567 dbname=dbdemo sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// app

	// db.DropTable(&User{}) // Delete old table if it exists.
	// db.CreateTable(&User{})

	// Save user list
	// for _, user := range users {
	//   db.Create(&user)
	// }

	//
	// user := User{Username: "sjobs"}
	db.Where(&User{Username: "sjobs"}).Delete(&User{})
	// fmt.Println(u)

	fmt.Println("Islem tmm.")
}

// Go object to PostgreSQL table
type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	{Username: "markzuck", FirstName: "Mark", LastName: "Zuck"},
	{Username: "bgates", FirstName: "Bill", LastName: "Gates"},
	{Username: "sjobs", FirstName: "Steve", LastName: "Jobs"},
	{Username: "lellison", FirstName: "Larry", LastName: "Ellison"},
}
