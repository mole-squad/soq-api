package focusareas

import (
	"github.com/burkel24/task-app/pkg/models"
	"github.com/go-chi/render"
)

func NewFocusAreaListResponseDTO(focusAreas []models.FocusArea) []render.Renderer {
	list := []render.Renderer{}
	for _, focusArea := range focusAreas {
		list = append(list, NewFocusAreaDTO(focusArea))
	}

	return list
}
