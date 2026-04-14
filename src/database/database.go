// Package database
package database

import (
	"fmt"

	"doko/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("doko.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := DB.AutoMigrate(&models.Book{}, &models.User{}); err != nil {
		panic(fmt.Errorf("error migrating database: %s", err.Error()))
	}
}
