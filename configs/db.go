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
	err := godotenv.Load()
	// "./configs/.env"
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())
		os.Exit(1)
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	// dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ")" + "/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"

	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(err.Error())
	}
	status := db.Find(&models.Status{})
	fmt.Println("Succesfully connected ")

	db.AutoMigrate(&models.Status{})
	db.AutoMigrate(&models.Task{})

	if status.RowsAffected == 0 {

		todo := models.Status{Status: "To do"}
		in_progress := models.Status{Status: "In progress"}
		done := models.Status{Status: "Done"}

		db.Create(&todo)
		db.Create(&in_progress)
		db.Create(&done)
	}
	return db
}
