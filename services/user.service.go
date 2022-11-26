package services

import (
	"github.com/AdrameDiakhate/golangTodoList.git/models"
	"gorm.io/gorm"
)

type IService interface {
	CreateUser(models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetOneUser(string) (models.User, error)
	VerifyEmail(string) (models.User, error)
	SayHello() string
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IService {
	return UserService{db: db}
}

// Create a new user

func (u UserService) CreateUser(user models.User) (createdUser models.User, err error) {
	u.db.Create(&user)
	return createdUser, u.db.Error
}

//Verify if an email exist in the database

func (u UserService) VerifyEmail(email string) (userFound models.User, err error) {
	var user models.User
	result := u.db.Table("users").Preload("Role").Where("email = ?", email).First(&user)
	return user, result.Error
}

//Get all users

func (u UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	return users, result.Error
}

//Get one user

func (u UserService) GetOneUser(IDUser string) (models.User, error) {
	var user models.User
	result := u.db.Where("ID=?", IDUser).First(&user)

	return user, result.Error
}

func (u UserService) SayHello() string {
	return "Adrame"
}
