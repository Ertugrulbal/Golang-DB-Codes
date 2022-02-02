// One to One

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

	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&User{})
	db.CreateTable(&User{})

	db.Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
		},
	})

	u := User{}
	c := Calendar{}
	db.First(&u).Related(&c, "calendar")

	fmt.Println(u)
	fmt.Println()
	fmt.Println(c)

	fmt.Println("Islem tmm.")
}

type User struct {
	gorm.Model
	Username   string
	FirstName  string
	LastName   string
	Calendar   Calendar
	CalendarID uint
}

type Calendar struct {
	gorm.Model
	Name string
}
