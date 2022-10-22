package services

import (
	"github.com/AdrameDiakhate/golangTodoList.git/configs"
	"github.com/AdrameDiakhate/golangTodoList.git/models"
)

// Create a new user

func CreateUser(user models.User) (createdUser models.User, err error) {
	db := configs.DBConnection()
	result := db.Create(&user)

	return createdUser, result.Error
}

//Verify if an email exist in the database

func VerifyEmail(email string) (userFound models.User, err error) {
	db := configs.DBConnection()
	var user models.User
	result := db.Table("users").Preload("Role").Where("email = ?", email).First(&user)
	return user, result.Error
}

//Get all users

func GetAllUsers() ([]models.User, error) {
	db := configs.DBConnection()
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

//Get one user

func GetOneUser(IDUser string) (models.User, error) {
	db := configs.DBConnection()
	var user models.User
	result := db.Where("ID=?", IDUser).First(&user)

	return user, result.Error
}
