package models

import (
	"gorm.io/gorm"
)

type TaskStatus int

const (
	TaskStatusOpen   TaskStatus = iota
	TaskStatusClosed            = iota
)

type Task struct {
	gorm.Model
	Summary string
	Notes   string
	Status  TaskStatus

	FocusAreaID uint
	FocusArea   FocusArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
