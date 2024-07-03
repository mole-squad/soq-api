package models

import (
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
	UserID  uint
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
