package main

import (
	"github.com/victorneuret/GitSync/server/src/app/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	db = database.ConnectDatabase()
	database.Boop()
}
