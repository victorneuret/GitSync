module github.com/victorneuret/GitSync/server/GitSync

go 1.12

require (
	github.com/jinzhu/gorm v1.9.2
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/victorneuret/GitSync/server/src/app/database v0.0.0-00010101000000-000000000000
)

replace github.com/victorneuret/GitSync/server/src/app/database => ./app/database
