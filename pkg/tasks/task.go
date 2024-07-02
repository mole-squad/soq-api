package tasks

import (
	"github.com/burkel24/task-app/pkg/users"
	"gorm.io/gorm"
)

type TaskStatus int

const (
	OpenStatus   TaskStatus = iota
	ClosedStatus            = iota
)

type Task struct {
	gorm.Model
	Summary string
	Notes   string
	Status  TaskStatus
	UserID  int
	User    users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
