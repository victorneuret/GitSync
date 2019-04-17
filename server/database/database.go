package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
}

func CloseDatabase() {
	_ = DB.Close()
}

func InitialMigration() {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(Repo{})
}