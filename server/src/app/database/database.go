package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db
}

func Boop() {
	fmt.Println("TEST")
}