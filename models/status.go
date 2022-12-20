package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ID     string `sql:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Status string `json:"status"`
}

func (u *Status) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("ID", uuid.New())
	return
}
