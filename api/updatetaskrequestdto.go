package api

import "net/http"

type UpdateTaskRequestDTO struct {
	Summary     string `json:"summary"`
	Notes       string `json:"notes"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *UpdateTaskRequestDTO) Bind(r *http.Request) error {
	return nil
}
