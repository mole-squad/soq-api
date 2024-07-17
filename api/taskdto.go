package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type TaskDTO struct {
	ID      uint   `json:"id"`
	Summary string `json:"summary"`
	Notes   string `json:"notes"`

	FocusArea FocusAreaDTO `json:"focusArea"`
}

func (t *TaskDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

var _ render.Renderer = &TaskDTO{}
