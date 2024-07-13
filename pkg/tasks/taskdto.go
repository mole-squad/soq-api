package tasks

import (
	"net/http"

	"github.com/mole-squad/soq/pkg/focusareas"
	"github.com/mole-squad/soq/pkg/models"
)

type TaskDTO struct {
	ID      uint   `json:"id"`
	Summary string `json:"summary"`
	Notes   string `json:"notes"`

	FocusArea focusareas.FocusAreaDTO `json:"focusArea"`
}

func NewTaskDTO(task models.Task) *TaskDTO {
	dto := &TaskDTO{
		ID:        task.ID,
		Summary:   task.Summary,
		Notes:     task.Notes,
		FocusArea: *focusareas.NewFocusAreaDTO(task.FocusArea),
	}

	return dto
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
