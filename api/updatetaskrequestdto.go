package api

import "net/http"

type UpdateTaskRequestDto struct {
	Summary     string `json:"summary"`
	Notes       string `json:"notes"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *UpdateTaskRequestDto) Bind(r *http.Request) error {
	return nil
}
