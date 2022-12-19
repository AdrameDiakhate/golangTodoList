package controllers

import (
	"github.com/AdrameDiakhate/golangTodoList.git/models"
	"github.com/AdrameDiakhate/golangTodoList.git/services"
	"github.com/gin-gonic/gin"
)

//Interface pour les controllers

type ITaskController interface {
	CreateTask(c *gin.Context)
	GetAllTask(c *gin.Context)
	GetOneTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type TaskController struct {
	service services.ITaskService
}

func NewTaskController(taskService services.ITaskService) ITaskController {
	return TaskController{service: taskService}
}

//Create a task

func (t TaskController) CreateTask(c *gin.Context) {
	task := models.Task{}

	err := c.ShouldBindJSON(&task)

	if err != nil {
		CreatedTask, err := t.service.CreateTask(task)
		if err != nil {
			c.JSON(201, gin.H{
				"message": "Task successfully created !",
				"data":    CreatedTask,
			})
		}
	}
	return
}

//Get all tasks
func (t TaskController) GetAllTask(c *gin.Context) {
	tasks, err := t.service.GetAllTask()
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Task successfully created !",
			"data":    tasks,
		})
	}

	c.JSON(200, gin.H{
		"message": "Task successfully created !",
		"data":    tasks,
	})

	// registers, err := r.Service.GetAllRegisters()
	// if err != nil {
	// 	setHttpResponse(c, http.StatusInternalServerError, failureResponse, nil, err)
	// 	return
	// }
	// setHttpResponse(c, http.StatusOK, successResponse, registers, err)
}

func (t TaskController) GetOneTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		CreatedTask, err := t.service.CreateTask(task)
		if err != nil {
			c.JSON(201, gin.H{
				"message": "Task successfully created !",
				"data":    CreatedTask,
			})
		}
	}
	return
}
func (t TaskController) UpdateTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		CreatedTask, err := t.service.CreateTask(task)
		if err != nil {
			c.JSON(201, gin.H{
				"message": "Task successfully created !",
				"data":    CreatedTask,
			})
		}
	}
	return
}

func (t TaskController) DeleteTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)

	if err != nil {
		CreatedTask, err := t.service.CreateTask(task)
		if err != nil {
			c.JSON(201, gin.H{
				"message": "Task successfully created !",
				"data":    CreatedTask,
			})
		}
	}
	return
}
