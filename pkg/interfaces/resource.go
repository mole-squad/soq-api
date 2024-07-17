package interfaces

import (
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/pkg/models"
)

type Resource interface {
	models.Model
	ToDTO() render.Renderer
}
