package tasks

import (
	"github.com/burkel24/task-app/pkg/models"
	"github.com/go-chi/render"
)

func NewTaskListResponseDTO(tasks []models.Task) []render.Renderer {
	list := []render.Renderer{}
	for _, task := range tasks {
		list = append(list, NewTaskDTO(task))
	}

	return list
}
