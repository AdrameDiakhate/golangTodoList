package routes

import (
	"github.com/AdrameDiakhate/golangTodoList.git/controllers"
	"github.com/gin-gonic/gin"
)

func setHelloRoute(r *gin.Engine) {
	hello := r.Group("/hello")
	var user controllers.UserController
	{
		hello.GET("/ada", user.SayHello)
	}
}
