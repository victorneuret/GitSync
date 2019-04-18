package database

import "github.com/jinzhu/gorm"

type Repo struct {
	gorm.Model
	Name           string
	Private        bool
	GithubURL      string
	Owner          string
	Updater        string
	UpdateOnMaster bool
}