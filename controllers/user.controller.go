package controllers

import (
	"net/http"

	"github.com/AdrameDiakhate/golangTodoList.git/models"
	"github.com/AdrameDiakhate/golangTodoList.git/services"
	"github.com/AdrameDiakhate/golangTodoList.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//Hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//Create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	user.Password = string(hash)
	user.ID = uuid.New().String()
	createdUser, error := services.CreateUser(user)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": createdUser.ID, "roleId": createdUser.RoleID, "email": createdUser.Email, "username": createdUser.Username})
}

func Login(c *gin.Context) {
	var authInfos utils.AuthInfos
	if err := c.ShouldBindJSON(&authInfos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFound, err := services.VerifyEmail(authInfos.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.AbortWithError(400, err)
		return
	}
	credentialError := userFound.CheckPassword(authInfos.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Les password ne correspondent pas"})
		c.Abort()
		return
	}
	tokenString, err := utils.GenerateJWT(userFound.Email, userFound.Username, userFound.Role.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

//Get all users
func GetAllUsers(c *gin.Context) {
	users, error := services.GetAllUsers()
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}

//Get one user

func GetOneUser(c *gin.Context) {
	IDUser := c.MustGet("ID").(string)
	//c.Params.ByName("ID")

	//c.MustGet("ID").(string)

	user, error := services.GetOneUser(IDUser)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})

}

func UpdateUser() {

}
func DeleteUser() {

}
