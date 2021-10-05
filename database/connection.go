package database

import (
	"demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Aa5986373@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}