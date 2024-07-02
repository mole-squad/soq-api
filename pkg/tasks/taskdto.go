package tasks

import "net/http"

type TaskDTO struct {
	Summary string `json:"summary"`
}

func NewTaskDTO(task Task) *TaskDTO {
	dto := &TaskDTO{Summary: task.Summary}

	return dto
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
