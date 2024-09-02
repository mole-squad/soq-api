package api

import (
	"fmt"
	"net/http"
)

type CreateFocusAreaRequestDTO struct {
	Name string `json:"name"`
}

func (dto *CreateFocusAreaRequestDTO) Bind(r *http.Request) error {
	if dto.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}
