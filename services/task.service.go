package services

import (
	"github.com/AdrameDiakhate/golangTodoList.git/models"
	"gorm.io/gorm"
)

//Interface pour les services

type ITaskService interface {
	CreateTask(models.Task) (models.Task, error)
	GetAllTask() ([]models.Task, error)
	GetOneTask(string) (models.Task, error)
	UpdateTask(string, models.Task) (models.Task, error)
	DeleteTask(string) (models.Task, error)
}

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) ITaskService {
	return TaskService{db: db}
}

//Create a task

func (t TaskService) CreateTask(task models.Task) (models.Task, error) {

	result := t.db.Create(&task)
	return task, result.Error

}

//Get all tasks

func (t TaskService) GetAllTask() ([]models.Task, error) {

	var tasks []models.Task
	err := t.db.Find(&tasks).Error
	return tasks, err

}

//Get a single task

func (t TaskService) GetOneTask(id string) (models.Task, error) {

	var task models.Task
	err := t.db.Where("id= ?", id).First(&task).Error
	return task, err

}

//Update a task

func (t TaskService) UpdateTask(id string, task models.Task) (models.Task, error) {

	result := t.db.Where("id =?", id).Updates(&task)
	return task, result.Error

}

//Delete a task

func (t TaskService) DeleteTask(id string) (models.Task, error) {

	var task models.Task
	result := t.db.Where("id = ?", id).Delete(&models.Task{})
	return task, result.Error

}
