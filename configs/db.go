package configs

import (
	"fmt"
	"log"
	"os"

	models "github.com/AdrameDiakhate/golangTodoList.git/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (db *gorm.DB) {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
		os.Exit(1)
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")" + "/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(err.Error())
	}
	fmt.Println("Succesfully connected ")

	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})

	//db.Create(&models.Role{RoleName: "admin"})
	//db.Create(&models.Role{RoleName: "user_simple"})

	return db
}
