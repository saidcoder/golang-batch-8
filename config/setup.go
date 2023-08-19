package config

import (
	"ujiKeterampilan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/api_uji_keterampilan?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Lakukan migrasi otomatis untuk tabel Book
	database.AutoMigrate(&models.Book{})
	// Lakukan migrasi otomatis untuk tabel Member
	database.AutoMigrate(&models.Member{})
	// Lakukan migrasi otomatis untuk tabel Loan
	database.AutoMigrate(&models.Loan{})

	DB = database
}
