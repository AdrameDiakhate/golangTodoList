package services

import (
	"github.com/AdrameDiakhate/golangTodoList.git/models"
	"gorm.io/gorm"
)

//Interface pour les services

type ITaskService interface {
	CreateTask(models.Task) (models.Task, error)
	GetAllTask() ([]models.Task, error)
	GetOneTask() (models.Task, error)
	UpdateTask(string) (models.Task, error)
	DeleteTask(string) (models.Task, error)
}

type TaskService struct {
	db *gorm.DB
}

func NewTaskService(db *gorm.DB) ITaskService {
	return TaskService{db: db}
}

func (t TaskService) CreateTask(task models.Task) (taskCreated models.Task, err error) {

	return taskCreated, err
}

func (t TaskService) GetAllTask() (tasks []models.Task, err error) {

	result := t.db.Find(&models.Task{})
	return tasks, result.Error
}

func (t TaskService) GetOneTask() (task models.Task, err error) {
	return task, err
}

func (t TaskService) UpdateTask(string) (task models.Task, err error) {
	return task, err
}

func (t TaskService) DeleteTask(string) (task models.Task, err error) {
	return task, err
}
