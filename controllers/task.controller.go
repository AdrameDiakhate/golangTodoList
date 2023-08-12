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

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(201, gin.H{
			"message": " Il y a une erreur!",
			"data":    err,
		})
	}

	CreatedTask, err := t.service.CreateTask(task)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Il y a une erreur!",
			"data":    err,
		})
	} else {

		c.JSON(201, gin.H{
			"message": "Task successfully created !",
			"data":    CreatedTask,
		})
	}
}

//Get all tasks

func (t TaskController) GetAllTask(c *gin.Context) {
	tasks, err := t.service.GetAllTask()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Il y a une erreur!",
			"data":    err,
		})
	}

	c.JSON(200, gin.H{
		"message": "All tasks !",
		"data":    tasks,
	})
}

//Get a single task

func (t TaskController) GetOneTask(c *gin.Context) {
	id := c.Params.ByName("id")

	task, err := t.service.GetOneTask(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Il y a une erreur!",
			"data":    err,
		})
	}

	c.JSON(200, gin.H{
		"message": "Single task!",
		"data":    task,
	})
}

//Update a task

func (t TaskController) UpdateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(500, gin.H{
			"message": " Il y a une erreur!",
			"data":    err,
		})
	}
	id := c.Params.ByName("id")

	updatedTask, err := t.service.UpdateTask(id, task)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erreur!",
			"data":    err,
		})
	}

	c.JSON(201, gin.H{
		"message": "Task successfully updated !",
		"data":    updatedTask,
	})
}

//Delete a task

func (t TaskController) DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")

	deletedTask, err := t.service.DeleteTask(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Il y a une erreur!",
			"data":    err,
		})
	} else {
		c.JSON(201, gin.H{
			"message": "Task successfully deleted !",
			"data":    deletedTask,
		})
	}
}



