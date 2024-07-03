package tasks

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/interfaces"
)

type TaskDTO struct {
	Summary string `json:"summary"`
}

func NewTaskDTO(task interfaces.Task) *TaskDTO {
	dto := &TaskDTO{Summary: task.Summary()}

	return dto
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
