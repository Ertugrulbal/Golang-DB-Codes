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
	db.CreateTable(&User{})

	fmt.Println("Islem tmm.")
}

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

func (u User) TableName() string {
	return "Userlar"
}
