package focusareas

import (
	"net/http"

	"github.com/burkel24/task-app/pkg/models"
	"github.com/go-chi/render"
)

type FocusAreaDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	TimeWindows []render.Renderer `json:"timeWindows"`
}

func NewFocusAreaDTO(focusArea models.FocusArea) *FocusAreaDTO {
	timeWindows := make([]render.Renderer, len(focusArea.TimeWindows))
	for i, timeWindow := range focusArea.TimeWindows {
		timeWindows[i] = NewTimeWindowDTO(timeWindow)
	}

	dto := &FocusAreaDTO{
		ID:          focusArea.ID,
		Name:        focusArea.Name,
		TimeWindows: timeWindows,
	}

	return dto
}

func (f *FocusAreaDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
