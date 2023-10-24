package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  "gin-gorm/helpers"
)

var DB *gorm.DB

func ConnectDatabase () {
  database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
  
  helpers.PanicError(err, "Opening database")

  err = database.AutoMigrate(Book{})

  helpers.PanicError(err, "Migrate database")

  DB = database
}
