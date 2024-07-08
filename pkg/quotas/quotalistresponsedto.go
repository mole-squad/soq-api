package quotas

import (
	"github.com/burkel24/task-app/pkg/models"
	"github.com/go-chi/render"
)

func NewQuotaListResponseDTO(tasks []models.Quota) []render.Renderer {
	list := []render.Renderer{}
	for _, task := range tasks {
		list = append(list, NewQuotaDTO(task))
	}

	return list
}
