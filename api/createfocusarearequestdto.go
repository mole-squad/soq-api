package api

import "net/http"

type CreateFocusAreaRequestDTO struct {
	Name string `json:"name"`
}

func (dto *CreateFocusAreaRequestDTO) Bind(r *http.Request) error {
	return nil
}
