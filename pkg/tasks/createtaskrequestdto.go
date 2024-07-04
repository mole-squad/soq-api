package tasks

import "net/http"

type CreateTaskRequestDto struct {
	Summary string `json:"summary"`
	Notes   string `json:"notes"`
}

func (dto *CreateTaskRequestDto) Bind(r *http.Request) error {
	return nil
}
