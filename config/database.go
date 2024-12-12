package config

import (
	"fmt"
	"os"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error pada .env file: %w", err)
	}
	return nil
}

func CreateDatabase() *gorm.DB {
	LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entities.User{}, entities.Admin{}, entities.Product{}, entities.Cart{}, entities.Order{}, entities.OrderDetail{}, entities.Profil{}, entities.Review{}, entities.Artikel{}, entities.Chat{})
	db.Migrator().CreateTable(&entities.Payment{})
	db.Migrator().CreateTable(&entities.Artikel{})
	db.Migrator().CreateTable(&entities.Chat{})
	db.Migrator().CreateTable(&entities.TotalPendapatan{})

	return db
}
