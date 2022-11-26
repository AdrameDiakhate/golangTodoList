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
	var UserRoute controllers.UserController
	{
		user.POST("/login", UserRoute.Login)
		//user.POST("/refresh",controllers.RefreshToken)
		secured.POST("/create", UserRoute.CreateUser)
		secured.GET("/all", UserRoute.GetAllUsers)
		secured.GET("/one", UserRoute.GetOneUser)
		// user.PUT("/update/:id", controllers.UpdateUser)
		// user.DELETE("/delete/:id", controllers.DeleteUser)
	}
}
