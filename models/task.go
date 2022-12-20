package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          string `sql:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	StatusID    string `gorm:"size:191" json:"statusID"`
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("ID", uuid.New())
	return
}
