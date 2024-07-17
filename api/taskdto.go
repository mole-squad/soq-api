package api

import (
	"net/http"
)

type TaskDTO struct {
	ID      uint   `json:"id"`
	Summary string `json:"summary"`
	Notes   string `json:"notes"`
	Status  int    `json:"status"`

	FocusArea FocusAreaDTO `json:"focusArea"`
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
