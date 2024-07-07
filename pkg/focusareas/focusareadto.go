package focusareas

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/models"
)

type FocusAreaDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewFocusAreaDTO(focusArea models.FocusArea) *FocusAreaDTO {
	dto := &FocusAreaDTO{
		ID:   focusArea.ID,
		Name: focusArea.Name,
	}

	return dto
}

func (f *FocusAreaDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
