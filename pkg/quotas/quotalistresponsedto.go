package quotas

import (
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/models"
)

func NewQuotaListResponseDTO(tasks []models.Quota) []render.Renderer {
	list := []render.Renderer{}
	for _, task := range tasks {
		list = append(list, NewQuotaDTO(task))
	}

	return list
}
