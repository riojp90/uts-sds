package database

import (
	"fmt"
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "username:password@tcp(your-database-host:port)/your-database-name?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		panic(err)
	}

	// AutoMigrate akan membuat tabel jika belum ada dan melakukan migrasi model User
	DB.AutoMigrate(&models.User{})
}
