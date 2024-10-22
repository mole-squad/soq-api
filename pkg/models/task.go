package models

import (
	"net/http"

	"github.com/burkel24/go-mochi"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"gorm.io/gorm"
)

type TaskStatus int

const (
	TaskStatusOpen TaskStatus = iota
	TaskStatusClosed
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

func (t *Task) GetID() uint {
	return t.ID
}

func (t *Task) ToDTO() render.Renderer {
	focusArea := t.FocusArea.ToDTO().(*api.FocusAreaDTO)

	dto := &api.TaskDTO{
		ID:        t.ID,
		Summary:   t.Summary,
		Notes:     t.Notes,
		Status:    int(t.Status),
		FocusArea: *focusArea,
	}

	return dto
}

func NewTaskFromCreateRequest(r *http.Request, user mochi.User) (*Task, error) {
	task := &Task{}

	dto := &api.CreateTaskRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	task.Summary = dto.Summary
	task.Notes = dto.Notes
	task.FocusAreaID = dto.FocusAreaID

	return task, nil
}

func NewTaskFromUpdateRequest(r *http.Request, user mochi.User) (*Task, error) {
	task := &Task{}

	dto := &api.UpdateTaskRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		return nil, err
	}

	task.Summary = dto.Summary
	task.Notes = dto.Notes
	task.FocusAreaID = dto.FocusAreaID

	return task, nil
}
