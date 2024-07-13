package api

import (
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/models"
)

func NewFocusAreaListResponseDTO(focusAreas []models.FocusArea) []render.Renderer {
	list := []render.Renderer{}
	for _, focusArea := range focusAreas {
		list = append(list, NewFocusAreaDTO(focusArea))
	}

	return list
}
