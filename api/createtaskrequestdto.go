package api

import "net/http"

type CreateTaskRequestDto struct {
	Summary     string `json:"summary"`
	Notes       string `json:"notes"`
	FocusAreaID uint   `json:"focusAreaId"`
}

func (dto *CreateTaskRequestDto) Bind(r *http.Request) error {
	return nil
}
