package routes

import (
	"github.com/AdrameDiakhate/golangTodoList.git/controllers"
	"github.com/AdrameDiakhate/golangTodoList.git/middleware"
	"github.com/gin-gonic/gin"
)

func setUserRoute(r *gin.Engine) {

	//User group: user

	user := r.Group("/user")
	secured := r.Group("/user/secured").Use(middleware.IsAuthenticated())
	{
		user.POST("/login", controllers.Login)
		secured.POST("/create", controllers.CreateUser)
		secured.GET("/all", controllers.GetAllUsers)
		secured.GET("/one", controllers.GetOneUser)
		// user.PUT("/update/:id", controllers.UpdateUser)
		// user.DELETE("/delete/:id", controllers.DeleteUser)
	}
}
