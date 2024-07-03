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
	summary string
	notes   string
	status  TaskStatus
	userID  uint
	user    users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Task) Summary() string {
	return t.summary
}

func (t *Task) Notes() string {
	return t.notes
}

func (t *Task) Status() TaskStatus {
	return t.status
}
