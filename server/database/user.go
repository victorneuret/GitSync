package database

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string
	Login     string
	Email     string
	AvatarURL string
	Token    string
}