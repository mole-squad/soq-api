package tasks

import "net/http"

type UpdateTaskRequestDto struct {
	Summary string `json:"summary"`
	Notes   string `json:"notes"`
}

func (dto *UpdateTaskRequestDto) Bind(r *http.Request) error {
	return nil
}
