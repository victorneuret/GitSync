package database

import (
	"fmt"
	"github.com/victorneuret/GitSync/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(config.Config.Database.Name, config.Config.Database.Parameter)
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