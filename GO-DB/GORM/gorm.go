package main

// go get -u github.com/jinzhu/gorm
// go get -u github.com/lib/pq

import (
	"github.com/jinzhu/gorm"
	// go get -u gorm.io/gorm
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=1234567 dbname=hr sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}

	println("Connection to database established.")
}
