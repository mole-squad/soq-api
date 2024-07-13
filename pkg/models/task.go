package models

import (
	"github.com/mole-squad/soq-api/api"
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

func (t *Task) AsDTO() *api.TaskDTO {
	focusArea := t.FocusArea.AsDTO()

	dto := &api.TaskDTO{
		ID:        t.ID,
		Summary:   t.Summary,
		Notes:     t.Notes,
		FocusArea: *focusArea,
	}

	return dto
}
