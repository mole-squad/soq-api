package api

import "net/http"

type UpdateFocusAreaRequestDTO struct {
	Name string `json:"name"`
}

func (dto *UpdateFocusAreaRequestDTO) Bind(r *http.Request) error {
	return nil
}
