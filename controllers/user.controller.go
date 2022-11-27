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

type IUser interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetOneUser(c *gin.Context)
	SayHello(c *gin.Context)
}

type UserController struct {
	userService services.IService
}

//Hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//Create a new user
func (u UserController) CreateUser(c *gin.Context) {
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
	createdUser, error := u.userService.CreateUser(user)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": createdUser.ID, "roleId": createdUser.RoleID, "email": createdUser.Email, "username": createdUser.Username})
}

func (u UserController) Login(c *gin.Context) {
	var authInfos utils.AuthInfos
	if err := c.ShouldBindJSON(&authInfos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFound, err := u.userService.VerifyEmail(authInfos.Email)

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
	tokenString, refreshToken, err := utils.GenerateJWT(userFound.Email, userFound.Username, userFound.Role.RoleName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "refreshToken": refreshToken})
}

// //Refresh token
// func RefreshToken(c gin.Context) (string, string, error) {
// 	type Refresh struct {
// 		Refresh string `json:"refresh"`
// 	}
// 	var refresh Refresh
// 	err := c.ShouldBind(&refresh)
// 	if err != nil {
// 		c.Errors.JSON()
// 	}
// 	return
// }

//Get all users
func (u UserController) GetAllUsers(c *gin.Context) {
	users, error := u.userService.GetAllUsers()
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

func (u UserController) GetOneUser(c *gin.Context) {
	IDUser := c.MustGet("ID").(string)
	//c.Params.ByName("ID")

	//c.MustGet("ID").(string)

	user, error := u.userService.GetOneUser(IDUser)
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

func (u UserController) SayHello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello Adramé Diakhaté DevOps",
	})
}
