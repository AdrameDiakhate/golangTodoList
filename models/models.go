package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       string `sql:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleName string
}

type User struct {
	gorm.Model
	ID       string `sql:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FistName string
	LastName string
	Email    string `gorm:"size:191; unique";`
	Username string `gorm:"size:191; unique";`
	Password string
	RoleID   string `gorm:"size:191"`
	Role     Role
	Task     []Task
}

type Task struct {
	gorm.Model
	ID          string `sql:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
	Status      string
	Deadline    *time.Time
	UserID      string
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("ID", uuid.New())
	return
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
