package api

import (
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/models"
)

func NewTaskListResponseDTO(tasks []models.Task) []render.Renderer {
	list := []render.Renderer{}
	for _, task := range tasks {
		list = append(list, NewTaskDTO(task))
	}

	return list
}
