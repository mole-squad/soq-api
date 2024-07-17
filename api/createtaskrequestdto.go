package api

import "net/http"

type CreateTaskRequestDTO struct {
	Summary     string `json:"summary"`
	Notes       string `json:"notes"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *CreateTaskRequestDTO) Bind(r *http.Request) error {
	return nil
}
