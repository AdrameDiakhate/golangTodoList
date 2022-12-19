package routes

import (
	"github.com/AdrameDiakhate/golangTodoList.git/controllers"
	"github.com/gin-gonic/gin"
)

type TaskRoute struct {
	Controller controllers.ITaskController
}

func NewTaskRoute(c controllers.ITaskController) TaskRoute {
	return TaskRoute{Controller: c}
}

func (taskRoute TaskRoute) setTaskRoute(r *gin.Engine) {

	// Task group: task
	task := r.Group("/task")
	{
		task.POST("/add", taskRoute.Controller.CreateTask)
		task.GET("/getalltasks", taskRoute.Controller.GetAllTask)
		task.GET("/getone/:id", taskRoute.Controller.GetOneTask)
		task.PUT("/update/:id", taskRoute.Controller.UpdateTask)
		task.DELETE("/delete/:id", taskRoute.Controller.DeleteTask)
	}
}
