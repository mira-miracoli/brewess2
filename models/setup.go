package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

        database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

        if err != nil {
                panic("Failed to connect to database!")
        }

        err = database.AutoMigrate(&Malt{},&Hop{},&Yeast{})
        if err != nil {
                return
        }

        DB = database
}