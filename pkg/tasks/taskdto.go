package tasks

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/models"
)

type TaskDTO struct {
	Summary string `json:"summary"`
	Notes   string `json:"notes"`
}

func NewTaskDTO(task models.Task) *TaskDTO {
	dto := &TaskDTO{
		Summary: task.Summary,
		Notes:   task.Notes,
	}

	return dto
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}