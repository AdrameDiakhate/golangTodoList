package routes

import (
	"github.com/AdrameDiakhate/golangTodoList.git/configs"
	"github.com/AdrameDiakhate/golangTodoList.git/controllers"
	"github.com/AdrameDiakhate/golangTodoList.git/services"
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine) {
	db := configs.DBConnection()
	taskService := services.NewTaskService(db)
	taskController := controllers.NewTaskController(taskService)

	taskRoute := NewTaskRoute(taskController)

	taskRoute.setTaskRoute(r)
}
